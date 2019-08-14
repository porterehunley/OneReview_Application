import 'bootstrap/dist/css/bootstrap.min.css';
import $ from 'jquery';
import Navbar from './navbar';
import Deck from "./Movies/deck";
import Popper from 'popper.js';
import 'bootstrap/dist/js/bootstrap.bundle.min';
import React, { Component } from "react";
import ReactDOM from 'react-dom';
import './index.css';

var moviesObj = [];

var moviesProm = new Promise((resolve, reject) => {
	let moviesList = [];
	$.ajax({
		url: 'https://cors-anywhere.herokuapp.com/http://truereview.network/api/movies/all',
		crossDomain: true,
		type: "GET",
		dataType: "json",
		success: function (result) {
			const moviesJson = result.movies;
			moviesJson.forEach(function(item, index) {
				let movie = new Object();
				movie.id = index;
				movie.header = "featured";
				movie.title = item["title"];
				movie.body = item["score"];
				moviesObj.push(movie);
			});
			resolve('success');
		},
		error: function () {
			resolve('error');
		}
	});
});

//Note that the class names and the elements it renders must be the same
class App extends Component {
	constructor(props){
		super(props);
	}
	render() {
		return(
			<div>
				<Navbar />
				<div className="cardContainer">
					<Deck movies={moviesObj}/>
				</div>
			</div>
		);
	}
}


//Render on the root div
moviesProm.then(function() {
	ReactDOM.render(
		<App />,
		document.getElementById('root')
	);
}).catch(function() {
	console.log("error");
});


