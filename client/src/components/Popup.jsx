import React, { useState } from "react";
import axios from "axios";

function Popup() {
  const [city, setCity] = useState("");

  const handleWeatherReport = (e) => {
    e.preventDefault();
    console.log(city)
    // alert("button clicked")
    const config = {
      headers: {
        "Access-Control-Allow-Origin": "*",
        "Access-Control-Allow-Methods": "GET,PUT,POST,DELETE,PATCH,OPTIONS"
      }
    };
    
    const response = axios.get("http://localhost:5000/getWeatherReport", config,  {
      city: city,
    });
    try {
      if (response) {
        console.log(response);

      }
    } catch (error) {
      console.log(error);
    }
  };
  return (
    <div>
      <button
        data-modal-target="authentication-modal"
        data-modal-toggle="authentication-modal"
        class="block text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
        type="button"
      >
        Weather Dashboard
      </button>

      <div
        id="authentication-modal"
        tabindex="-1"
        aria-hidden="true"
        class="hidden overflow-y-auto overflow-x-hidden fixed top-0 right-0 left-0 z-50 justify-center items-center w-full md:inset-0 h-[calc(100%-1rem)] max-h-full"
      >
        <div class="relative p-4 w-full max-w-md max-h-full">
          <div class="relative bg-white rounded-lg shadow-sm dark:bg-gray-700">
            <div class="flex items-center justify-between p-4 md:p-5 border-b rounded-t dark:border-gray-600 border-gray-200">
              <h3 class="text-xl font-semibold text-gray-900 dark:text-white">
                Enter yout city name to get the weather
              </h3>
            </div>

            <div class="p-4 md:p-5">
              <form class="space-y-4" action="#">
                <div>
                  <label
                    for="text"
                    class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                  >
                    City
                  </label>
                  <input
                    type="text"
                    name="city"
                    id="city"
                    onChange={(e) => setCity(e.target.value)}
                    class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white"
                    placeholder="e.g.Mathura"
                    required
                  />
                </div>
                <button
                  type="submit"
                  onClick={handleWeatherReport}
                  class="w-full text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
                >
                  Get Weather report
                </button>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
export default Popup;
