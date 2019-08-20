import React from 'react';

//This is going to be a functional component because its state is taken from the deck
//This just returns a view
//TODO: Make these clickable
function Movie(props) {
	return(
		<div className="card">
			<div className="card-header">
			  {props.header}
			</div>
		  <img className="card-img-top" src={props.image} alt="Card image cap"></img>
		  <div className="card-body">
		    <h5 className="card-title">{props.title}</h5>
		    <p className="card-text">{props.body}</p>
		    <a href="#" className="btn btn-primary">Movie Page</a>
		  </div>
		</div>
	);
}
		



export default Movie;