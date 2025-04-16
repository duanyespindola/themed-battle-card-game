# REQ-B85E - Waiting Room Status

A waiting room is created with one player inside and has the status "Waiting for Another Player"
When a secod player enters the room, the status changes to "Waiting for the match starts"
If a player leaves the room before the match starts, the status turns back to "Waiting for another player"
If all players leave the room, the status changes to "Empty"
When a match starts the status changes to "Match in progress"
When a match is finished the status changes to "Match is finished"