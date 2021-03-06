/*
 * File: /server/league.go                                                     *
 * Project: tdd                                                                *
 * Created At: Friday, 2022/06/24 , 06:32:56                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Friday, 2022/06/24 , 11:57:33                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package server

import (
	"encoding/json"
	"go_start/tdd/model"
	"net/http"
)

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")

	json.NewEncoder(w).Encode(p.Store.GetLeague())
}

func (p *PlayerServer) getLeagueTable() model.League {
	return model.League{
		{"Chris", 20},
	}
}
