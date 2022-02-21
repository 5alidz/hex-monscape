# Project Methodology

This document contains how-to guidelines to create web app project using Hexagonal Architecture.

1. Understand business requirements.
2. Validate your understanding by creating document that contains high level overview of the application. Usually we use `README.md` for this.
3. Create use case list for the app. This list better accompanied by use case diagrams.
4. Create list of the API based on the use case list. Write it on `docs/http_api.md`. By creating the API documentation first, it will be easier for you to implement the API because everything already clear from the beginning.
5. Find the usage context of your API, this is to spot out your core components.
6. Create class diagram for the core components. By creating the class diagram, it will be easier for you to structure your code. Use [UMLet](https://www.umlet.com/) to create the class diagram.
7. Finish up the class diagram until you satisfied enough.
8. Implement your application, make any necessary adjustment to the docs.
9. Create test for your application.
10. Dockerize your application.