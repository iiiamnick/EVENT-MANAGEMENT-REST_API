package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iiiamnick/GOLANG--REST-API-EVENT-BOOKING-.git/models"
)

func getEvents(context *gin.Context) {
	events, err := models.GetALLEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the data"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the requested data"})
		return

	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the event"})
		return
	}
	context.JSON(http.StatusOK, event)

}

func createEvents(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not process the request"})
	}

	userId := context.GetInt64("userId")
	event.UserID = userId
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not return the event, Try Again Later!!!"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "status created!", "event": event})

}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the requested data"})
		return

	}
	userId := context.GetInt64("userId")

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "couldnt fetch the event"})
		return
	}
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized to update the event"})
		return
	}

	var updateEvent models.Event
	err = context.ShouldBindJSON(&updateEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the requested data"})
		return
	}
	updateEvent.ID = eventId
	err = updateEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "couldnt update the event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated succesfully"})

}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the event"})
		return
	}
	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "coudnt fetch the data"})
		return
	}
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized to delete the event"})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "coudnt delete the event"})
		return

	}

	context.JSON(http.StatusOK, gin.H{"message": "Event successfully deleted"})

}
