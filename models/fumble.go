package models

import (
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/TF2Stadium/Helen/config"
	"github.com/TF2Stadium/fumble/mumble"
)

func FumbleLobbyCreated(lob *Lobby) error {
	if config.Constants.FumblePort == "" {
		return nil
	}

	err := call(config.Constants.FumblePort, "Fumble.CreateLobby", lob.ID, &struct{}{})

	if err != nil {
		logrus.Warning(err.Error())
		return err
	}

	return nil
}

func fumbleAllowPlayer(lobbyId uint, playerName string, playerTeam string) error {
	if config.Constants.FumblePort == "" {
		return nil
	}

	user := mumble.User{}
	user.Name = playerName
	user.Team = mumble.Team(playerTeam)

	err := call(config.Constants.FumblePort, "Fumble.AllowPlayer", &mumble.LobbyArgs{user, lobbyId}, &struct{}{})
	if err != nil {
		logrus.Warning(err.Error())
	}

	return nil
}

func FumbleLobbyPlayerJoinedSub(lob *Lobby, player *Player, slot int) {
	if config.Constants.FumblePort == "" {
		return
	}

	team, class, _ := LobbyGetSlotInfoString(lob.Type, slot)
	fumbleAllowPlayer(lob.ID, strings.ToUpper(team)+"_"+strings.ToUpper(class), strings.ToUpper(team))
}

func FumbleLobbyPlayerJoined(lob *Lobby, player *Player, slot int) {
	if config.Constants.FumblePort == "" {
		return
	}

	team, class, _ := LobbyGetSlotInfoString(lob.Type, slot)
	fumbleAllowPlayer(lob.ID, strings.ToUpper(team)+"_"+strings.ToUpper(class), "")
}

func FumbleLobbyEnded(lob *Lobby) {
	if config.Constants.FumblePort == "" {
		return
	}

	err := call(config.Constants.FumblePort, "Fumble.EndLobby", lob.ID, nil)
	if err != nil {
		logrus.Warning(err.Error())
	}
}
