package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sebdah/recharged-admin/models"
	"gopkg.in/mgo.v2"
)

// Create new BootNotificationLog
func BootNotificationLogCreateHandler(w http.ResponseWriter, r *http.Request) {
	bootNotificationLog := models.NewBootNotificationLog()
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&bootNotificationLog)
	if err != nil {
		log.Debug("Unable to parse request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check required fields
	if bootNotificationLog.Model == "" {
		log.Debug("Missing required field: model")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if bootNotificationLog.Vendor == "" {
		log.Debug("Missing required field: vendor")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Save the object
	err = models.Save(bootNotificationLog)
	if err != nil {
		if mgo.IsDup(err) {
			w.WriteHeader(http.StatusConflict)
			log.Debug("BootNotificationLog already exists")
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			log.Error("Error in MongoDB communication: %s", err.Error())
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	js, _ := json.Marshal(bootNotificationLog)
	w.Write(js)

	return
}
