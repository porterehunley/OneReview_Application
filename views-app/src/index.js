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
			for (let i = 0; i < moviesJson.length; i++) {
				moviesList.push(moviesJson[i]["title"]);
			}
			moviesList.forEach(function(item, index) {
				let movie = new Object();
				movie.id = index;
				movie.header = "featured";
				movie.title = item;
				movie.body = "succ";
				moviesObj.push(movie);
			});
			resolve('success');
		},
		error: function () {
			resolve('error');
		}
	});
});

const movies = [{
	id: 1,
	header: "featured",
	title: "Shark boy",
	body: "This is a mediocre movie",
}, 
{
	id: 2,
	header: "featured",
	title: "Lava Girl",
	body: "This is a better movie",

}, 
{
	id: 3,
	header: "featured",
	title: "Lava Girl",
	body: "This is a better movie",
},
{
	id: 4,
	header: "featured",
	title: "Lava Girl",
	body: "This is a better movie",
},
{
	id: 5,
	header: "featured",
	title: "Lava Girl",
	body: "This is a better movie, Now that we covered all the “easy” things let’s move into the “real world debugging” — working with breakpoints. A breakpoint is an instruction in code which will stop all the execution and fire up a debugger. There are 2 ways of setting up breakpoints that you will find handy",
},
{
	id: 6,
	header: "featured",
	title: "Lava Girl",
	body: "This is a better movie",
},
{
	id: 7,
	header: "featured",
	title: "Lava Girl",
	body: "This is a better movie",
}]


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
	console.log("succ");
	ReactDOM.render(
		<App />,
		document.getElementById('root')
	);
}).catch(function() {
	console.log("error");
});


