import React from "react";
import { Provider } from "react-redux";
import { BrowserRouter as Router } from "react-router-dom";

import { CustomersRoute } from "./features/customers";

import { store } from "./app/store";

const App = () => {
  return (
    <Provider store={store}>
      <Router>
        <CustomersRoute />
      </Router>
    </Provider>
  );
};

export default App;
