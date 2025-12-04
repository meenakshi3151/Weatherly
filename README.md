# Weatherly

This project is a full-stack weather reporting application built using **Go (Golang)** for the backend API and **React.js** for the frontend interface. The application allows users to enter a city name and fetch real-time weather details using external weather and geocoding APIs.

---

## Overview

The backend service fetches latitude and longitude for a given city and then retrieves weather details based on those coordinates. The frontend provides a clean UI that displays weather information through reusable components such as **Header**, **Popup**, and **Footer**.

---

## Tech Stack

### **Backend (Go)**
- Go (Golang)
- Gorilla Mux (Routing)
- CORS Middleware
- godotenv for environment variables
- External Geocoding API (City → Coordinates)
- External Weather API (Lat/Lon → Weather Data)

### **Frontend (React)**
- React.js
- Reusable components (Header, Popup, Footer)

---

## Features

- Fetches live weather data for any entered city  
- Converts city names to coordinates using Geocoding API  
- Calls Weather API using latitude and longitude  
- CORS-enabled backend compatible with React frontend  
- Responsive frontend interface  
- Clean and modular codebase  

---

## 🔧 Environment Variables

The backend uses an `.env` file containing:  
- `API_KEY` → API key for weather/geocoding services  
- `GEOENCODING_API` → URL for city-to-coordinates API  
- `WEATHER_API` → URL for weather data API  

---

## Endpoints

**GET /** → Returns basic server message  
**GET /hello** → Test endpoint  
**GET /getWeatherReport?city=<city_name>** → Returns weather data for the given city  

---

## Project Purpose

This project demonstrates how to integrate a **Go REST API** with a **React client**, handle external API calls, manage CORS, and build a functional weather reporting interface.

---

## Contributions

Feel free to contribute enhancements, UI improvements, or optimizations to the backend.

---

## License

This project is open-source and available under the MIT License.
