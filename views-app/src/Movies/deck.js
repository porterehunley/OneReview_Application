import React from "react";
import Movie from "./movie";

const Deck = ({movies}) => {
	//Deck controls the Array of movies
	//Number of movies is handled by the App
	//Titles of the movies are also handled by the App for search functionality
	//Deck is in charge of ordering them?? NO!
	//To render an array of movies,
	const moviesArray = movies.map(movie => (
		<Movie key={movie.id}
			header={movie.header}
			title={movie.title}
			body={movie.body} 
			image={movie.image}/>
	));

	return (
		<div className="card-deck">
			{moviesArray}
		</div>
	);

	
	//OLD CODE BELOW COMMENT OUT AT OWN RISK
	// constructor(props) {
	// 	super(props);
	// 	this.state = {
	// 		movies: Array(this.props.numberOfMovies).fill(null),
	// 	};
	// }

	// //This function should set the index of an array to the movie 
	// setMovie(i) {
	// 	const movies = this.state.movies.slice();
	// 	movies[i] = {title: this.props.titles[i],
	// 							header: this.props.headers[i],
	// 							body: this.props.bodies[i]};
	// 	this.setState({
	// 		movies: movies,
	// 	});
	// }





	// render() {
	// 	return(
	// 		<div class="card-deck">
	// 			<Movie movies={this.state.movies}/>
	// 		</div>
	// 	);
	// }
};

export default Deck;