import 'bootstrap/dist/css/bootstrap.min.css';
import $ from 'jquery';
import Navbar from './navbar';
import Popper from 'popper.js';
import 'bootstrap/dist/js/bootstrap.bundle.min';
import React, { Component } from "react";
import ReactDOM from 'react-dom';
import './index.css';

//Note that the class names and the elements it renders must be the same
class App extends Component {
	render() {
		return(
			<p>Hello Appp</p>
		);
	}

}


//Render on the root div
ReactDOM.render(
	<App />,
	document.getElementById('root')
);

ReactDOM.render(
  <Navbar />,
  document.getElementById('nav')
);