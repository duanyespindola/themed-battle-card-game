# REQ-B85E - Waiting Room Status

A waiting room is created with one player inside and has the status "Waiting for Another Player"
When a secod player enters the room, the status changes to "Waiting for the match starts"
If a player leaves the room before the match starts, the status turns back to "Waiting for another player"
If all players leave the room, the status changes to "Empty"
When a match starts the status changes to "Match in progress"
When a match is finished the status changes to "Match is finished"


# REQ-E973 - Room ID

The room ID should be a UUID with 36 positions, including the dashs.
Example: "08a1039e-8dc4-4379-a2f2-15f166bd315d"

# REQ-9C5C - Player ID

The player ID should be a UUID with 36 positions, including the dashs.
Example: "08a1039e-8dc4-4379-a2f2-15f166bd315d"
