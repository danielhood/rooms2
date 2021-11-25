# Rooms II
Web UI version of the Rooms text adventure supporting multiple concurrent users.  The front-end is an Angular web app and the backend is a set of rest services implemented in Go.

## Goals

The goal of this project, other than an exercise in Angular and Go, is to create a simple platform to help kids explore basic adventure/role based game design and simple programming. A text based adventure was chosen to allow game concepts to be quickly implemented, encouraging reading/writing, and use of imagination over graphics.

The current base could allow for incorporation of graphic elements for extending the web UI (icon design, etc), character images, and could be extended to include _static_ graphics for rooms (graphic adventure style).

## What's Done

- Login/Auth
- UI and service layer for sending commands and displaying results

## What's Next

- Command processing (importing much of the original Rooms project)
- Introduce an event model for server events (hit by another character, another character walks into the room, user messages, etc)
- UI and backend model for user stats (health, location, inventory)
- Model for user properties (race, class, etc) if we decide to go with an rpg style adventure as opposed to just a human wandering and interacting with a world
- Persist user state and world state
- world gen (auto gen or designed?)
- Introduce multi-user support and interactions
-- PvP (turn based?), Messaging
- Server ticks independent of user interaction (autonomous objects)
- NPCs
- wold builder through web interface
