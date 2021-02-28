/* global fetch */
require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");

let state = {};

// loadAttributeData is used on the edit criteria page, to populate the page with all current attribute data, as well as
// all the different attributes that are in the membership database
const loadAttributeData = function(attrs, initial, groupCriteriaID) {
  state.attributes = attrs;
  state.criteria = initial;

  for (const [key, values] of Object.entries(state.criteria)) {
    for (let i = 0; i < values.length; i++) {
      const splitValue = values[i].split(":");
      let value;
      let operator;
      if (splitValue.length == 2) {
        value = splitValue[1];
        switch(splitValue[0]) {
          case "lt": operator = "<"; break;
          case "gt": operator = ">"; break;
          case "lte": operator = "<="; break;
          case "gte": operator = ">="; break;
          case "ne": operator = "!="; break;
          default: operator = "=";
        }
      } else {
        value = values[i];
        operator = "=";
      }
      $(`#${groupCriteriaID}`).append(
        `<span class="badge badge-secondary badge-criteria">
          ${key} ${operator} ${value} <i class="far fa-times-circle delete-attr" title="Remove criterion"
          onclick="App.removeCriteria(this.parentElement, '${key}', '${value}')"></i></span> `);
    }
  }

  let names = Object.entries(attrs);
  names.sort(function(x, y) { return x[1].description > y[1].description; });
  for (let [key, value] of names) {
    $('#newCriteriaName').append($('<option>', {
      value: key,
      text: value.description
    }));
  }
};

// updateCriteriaValues is called when the "name" dropdown is changed on the edit criteria page. This updates the values
// dropdown on the same edit criteria page to reflect each attribute name
const updateCriteriaValues = function(select, valueID) {
  $(valueID).empty();
  if (select.value == "_blank") {
    return;
  }
  attr = state.attributes[select.value];
  for (let i = 0; i < attr.values.length; i++) {
    value = attr.values[i];
    $(valueID).append($('<option>', {
      value: value,
      text: value
    }));
  }
};

// addCriteria is triggered when the add button is pressed on the edit criteria page. This adds attribute criteria to
// the current criteria locally
const addCriteria = function(attrID, attrOp, attrValue, currentCriteriaID) {
  const attr = document.getElementById(attrID).value;
  const displayValue = document.getElementById(attrValue).value;
  const operator = document.getElementById(attrOp).value;
  let dbValue;
  let displayOperator;

  switch(operator) {
    case "lt": displayOperator = "<"; dbValue = [operator, displayValue].join(":"); break;
    case "gt": displayOperator = ">";  dbValue = [operator, displayValue].join(":");break;
    case "lte": displayOperator = "<=";  dbValue = [operator, displayValue].join(":");break;
    case "gte": displayOperator = ">=";  dbValue = [operator, displayValue].join(":");break;
    case "ne": displayOperator = "!=";  dbValue = [operator, displayValue].join(":");break;
    default: displayOperator = "="; dbValue = displayValue;
  }

  let added = false;
  if (state.criteria) {
    if (!state.criteria[attr] || state.criteria[attr].length == 0) {
      state.criteria[attr] = [dbValue];
      added = true;
    } else if (state.criteria[attr].indexOf(displayValue) < 0) {
      state.criteria[attr] = state.criteria[attr].concat(dbValue);
      added = true;
    }
  }

  if (added) {
    $(`#${currentCriteriaID}`).append(
    `<span class="badge badge-secondary badge-criteria">
      ${attr} ${displayOperator} ${displayValue} <i class="far fa-times-circle"
      onclick="App.removeCriteria(this.parentElement, '${attr}', '${dbValue}')"></i></span>`);
  }
};

// removeCriteria is triggered when the user clicks the "x" icon on the edit criteria page, removing it from the
// current criteria
const removeCriteria = function(element, attr, value) {
  if (state.criteria[attr]) {
    element.remove();
    state.criteria[attr] = state.criteria[attr].filter(function(v, index, arr) { v != value; });
    if (state.criteria[attr].length == 0) {
      delete state.criteria[attr];
    }
  }
};

// submit commits the changes to a group's criteria to the backend
const submit = function(id) {
  fetch('/groups/'+id+'/criteria/edit', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(state.criteria),
  }).then(resp => {
    if (resp.status >= 200 && resp.status < 400) {
      if (resp.redirected) {
        document.location.href = resp.url;
      }
    } else {
      throw Error(resp.statusText);
    }
  }).catch(err => {
    alert(err);
  });
};

// Adds filtering to lists
const filterList = function(e, listId) {
  let input, filterTerm, listEl, listItems, li, contentEl, txtVal;
  input = e.target;
  filterTerm = input.value.toUpperCase();
  listEl = $(`#${listId}`);
  listItems = listEl.find('li');

  listItems.each(i => {
    li = $(listItems.get(i));
    contentEl = li.find('.list-item-content');
    txtVal = $(contentEl).text();

    if (txtVal.toUpperCase().includes(filterTerm)) {
      li.show();
    } else {
      li.hide();
    }
  });
};

const defaultFetchParams = {
  headers: { 'Content-Type': 'application/json' },
  credentials: 'same-origin',
};
const simpleFetch = async (url, params) => {
  const result = await fetch(url, { ...params, ...defaultFetchParams });
  if (!result.ok) {
    throw new Error(result.status, result.statusText);
  }
  return result.json();
};

function isValidEmail(email) {
  // eslint-disable-next-line no-useless-escape
  const regex = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
  return regex.test(String(email).toLowerCase());
}

const searchUser = async () => {
  const email = $(`#search-user`)[0].value
  if (!isValidEmail(email)) {
    alert('please submit a valid email');
    return
  }

  try {
  await simpleFetch(`/members/${email}/groups/`,{method:'GET'});
}

  catch(e) {
    alert('sorry - something went wrong trying to find the user.')
    return
  }
  window.location.pathname = `members/${email}/groups`;
}

const filterTable = function(e, tableId) {
  let input, filterTerm, tableEl, tbodyEl, trItems, tr, tdItems, trShow, td, txtVal;
  input = e.target;
  filterTerm = input.value.toUpperCase();
  tableEl = $(`#${tableId}`);
  tbodyEl = $(tableEl.find('tbody'));
  trItems = tbodyEl.find('tr');

  trItems.each(i => {
    tr = $(trItems.get(i));
    tdItems = tr.find('.table-filterable');

    trShow = false;
    tdItems.each(j => {
      td = $(tdItems.get(j));
      txtVal = td.text()
      if (txtVal.toUpperCase().includes(filterTerm)) {
        trShow = true
      }
    });

    if (trShow) {
      tr.show();
    } else {
      tr.hide();
    }
  });
};

// Need to expose these functions under the App global object
module.exports = {
  loadAttributeData,
  updateCriteriaValues,
  addCriteria,
  removeCriteria,
  submit,
  filterList,
  searchUser,
  filterTable,
};
