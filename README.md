# Youtube Review Application 

This application creates a score for a movie from its response on YouTube. "Response" is a combination of the scores given by review videos, and how well those review videos were received.

## Backend Breakdown

This project is composed of three seperate applications running in paralell. The chart below should show how everything interacts.

![Backend Flowchart](https://i.ibb.co/2vsQKks/Screen-Shot-2019-08-21-at-12-03-30-PM.png)

### Frontend App

The front-end application is built using **React** and hosted on Google's App Engine. 

The design of the front-end app is centered around display; it handles almost no logic. All of the pcitures, titles, scores, are fetched from the Golang backend via a RESTful API. 

##### Frameworks Used

The front-end uses **Bootstrap** for style. It specifically makes use of the fantastic Bootstrap Cards.



### Backend for Frontend App

The backend for frontend is written in **Golang** and utilizes a **Kubernetes** deployment.

This application is made to handle anything that the frontend might need. Some things include the collection of pcitures for the movies, serving pre-proccessed data like titles, scores, descriptions, etc.

The application also acts as a wrapper for the NoSQL Client database. This allows for other applications to communicate with the database via a RESTful API. 

In the future, this application will handle user logins and authentication. 

This application uses nothing outside of the standard Go library.



### Data-Side App

The data-side application collects, cleans, and maintains the data nesscesary to develop the models used to determine review scores. 

Here is an overview of the pipeline and application. 

![Data Pipeline](https://i.ibb.co/MBgh5gS/Screen-Shot-2019-08-21-at-1-03-50-PM.png)

##### Deployment

This application uses **CircleCI** to package/deploy to the Python Package Index (PyPI). **Ansible** provisions and deploys the package to a linux server. 

##### GUI??

Yes, there is a whole GUI for this side of the appliation which can be found [here](https://truereview.dev). If you are interested in seeing it in action, email me at porterhunley@gatech.edu and I will send you the credentials. 

The GUI acts as an interface to the data collection process. It allows you to add/remove movies from the database. It also shows you which data might not be useful and needs to be looked at. In addition, it acts as an easy interface to manually input the "scores" for the video reviews.

##### Frameworks Used

This app uses **Flask** for the back-end, and **JQuery** for the front.


### Is Everything Working??

Absolutely not. I am still honing in the model into something I would be comfortable deploying. In addition, the data-pipeline does not clean the data as well as it should. 












