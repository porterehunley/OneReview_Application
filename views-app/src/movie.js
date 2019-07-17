import React, { Component } from 'react';

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