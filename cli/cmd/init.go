package cmd

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/GracepointMinistries/hub/client"
	oidc "github.com/coreos/go-oidc"
	"github.com/fatih/color"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

var (
	host          string
	defaultHost   string
	authConfig    *oauth2.Config
	codeChallenge string
	server        *http.Server
	clientID      = "582682625807-uginu295mmv5o5kqd8v03eqv3sfcrild.apps.googleusercontent.com"
	clientSecret  = "yedyO6TKEpwXblhldQXq1C7q"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the CLI",
	Run: func(cmd *cobra.Command, args []string) {
		if host == "" {
			fmt.Fprintln(os.Stderr, `Error: required flag(s) "host" not set`)
			fmt.Fprintln(os.Stderr, cmd.UsageString())
			os.Exit(1)
		}
		runInit()
	},
}

func runInit() {
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, "https://accounts.google.com")
	checkError(err)
	authConfig = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  "http://localhost:8080/callback",
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}
	codeChallenge = generateChallenge(32)
	challenge := hash(codeChallenge)
	url := authConfig.AuthCodeURL(
		"state",
		oauth2.SetAuthURLParam("code_challenge", challenge),
		oauth2.SetAuthURLParam("code_challenge_method", "S256"),
	)
	fmt.Println(color.CyanString("You will now be taken to your browser for authentication"))
	time.Sleep(1 * time.Second)
	open.Run(url)
	time.Sleep(1 * time.Second)
	fmt.Printf("Authentication URL: %s\n", url)

	mux := http.NewServeMux()
	mux.HandleFunc("/callback", callbackHandler)
	server = &http.Server{Addr: ":8080", Handler: mux}
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		checkError(err)
	}
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	queryParts, err := url.ParseQuery(r.URL.RawQuery)
	checkError(err)

	code := queryParts["code"][0]
	token, err := authConfig.Exchange(
		r.Context(),
		code,
		oauth2.SetAuthURLParam("code_verifier", codeChallenge),
		oauth2.SetAuthURLParam("client_id", clientID),
	)
	checkError(err)

	c := client.NewAPIClient(client.NewConfiguration())
	c.ChangeBasePath(host)

	sessionToken, _, err := c.AuthApi.ExchangeAdmin(r.Context(), client.TokenPayload{
		Token: token.AccessToken,
	})
	checkError(err)
	fileConfig.Host = host
	fileConfig.Token = sessionToken.Token
	writeConfigFile()

	msg := "<p><strong>Success!</strong></p>"
	msg = msg + "<p>You are authenticated and can now return to the CLI.</p>"
	fmt.Fprintf(w, msg)
}

func generateChallenge(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length, length)
	for i := 0; i < length; i++ {
		b[i] = byte(r.Intn(255))
	}
	return encode(b)
}

func hash(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	return encode(h.Sum(nil))
}

func encode(msg []byte) string {
	encoded := base64.StdEncoding.EncodeToString(msg)
	encoded = strings.Replace(encoded, "+", "-", -1)
	encoded = strings.Replace(encoded, "/", "_", -1)
	encoded = strings.Replace(encoded, "=", "", -1)
	return encoded
}

func init() {
	initCmd.Flags().StringVarP(&host, "host", "H", defaultHost, "host to use for initialization")

	rootCmd.AddCommand(initCmd)
}
