import 'bootstrap/dist/css/bootstrap.min.css';
import $ from 'jquery';
import Popper from 'popper.js';
import 'bootstrap/dist/js/bootstrap.bundle.min';
import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';

//Note that the class names and the elements it renders must be the same
class Movie extends React.Component {
	render() {
		return (
			<div class="card">
				<div class="card-header">
				  Featured
				</div>
			  <img class="card-img-left" src="https://cdn.images.express.co.uk/img/dynamic/36/590x/1153955_1.jpg" alt="Card image cap"></img>
			  <div class="card-body">
			    <h5 class="card-title">Card title</h5>
			    <p class="card-text">Some quick example text to build on the card title and make up the bulk of the card's content.</p>
			    <a href="#" class="btn btn-primary">Go somewhere</a>
			  </div>
			</div>
		);
	}
}


//Render on the root div
ReactDOM.render(
	<app />,
	document.getElementById('root')
);