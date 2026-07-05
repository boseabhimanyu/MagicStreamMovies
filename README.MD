# MagicStream 🎬✨

Movie streaming platform with AI recommendation built with modern web technologies (React/Go/gin-gonic/MongoDB) 

---

## About  

This project is a full-stack simulation of a modern **Movie Streaming Platform**, designed to showcase how different technologies can be combined to deliver a scalable, AI-powered application.  

The system brings together a **React-based frontend** for an engaging user experience, a **Go-based backend** for high-performance API services that runs on the gin (gin-gonic) web framework, and an **AI-powered recommendation engine** to personalize movie suggestions using **LangChainGo** and **OpenAI**.  

It also demonstrates how **MongoDB** can serve as a reliable, scalable database solution for managing media metadata and user preferences.  

---

## Features

- Movie Streaming service simulated on the front end using React and React-Player
- Web API service written using GO and runs on the gin-gonic web framework 
- AI Recommendation service using LangChainGo, Go and OpenAI
- Scalable backend storage provided by MongoDB

---

## Tech Stack

| Frontend / Client | JavaScript / React |
| Backend / Server | Go / gin-gonic |
| Storage / Database | MongoDB |
 
---

## Link to Video Tutorial on How to Build the App
- https://youtu.be/jBf7of9JTV8

## Environment Variables Server

To run this project, you will need to add the following environment variables to your .env file

```
DATABASE_NAME=<db_name>

MONGODB_URI=<mongo_uri>

SECRET_KEY=<any_randon_text_string>

SECRET_REFRESH_KEY=<any_randon_text_string>

BASE_PROMPT_TEMPLATE=Return a response using one of these words: {rankings}. The reposne should be a single word and should not contain any other text. The response should be based on the following review:

OPENAI_API_KEY=<open_api_key>

RECOMMENDED_MOVIE_LIMIT=5

ALLOWED_ORIGINS=http://localhost:3000,http://localhost:5173,http://localhost:8080
```