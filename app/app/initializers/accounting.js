import { currency, number } from "accounting/settings";

export default {
  name: 'accounting.js',
  initialize: function() {
    currency.symbol = "€";
    currency.format = "%v %s";
    currency.decimal = ",";
    currency.thousand = ".";
    currency.precision = "0";
  }
};
