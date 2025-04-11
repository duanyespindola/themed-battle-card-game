Feature: Waiting Room

    Scenario: First player joins the game
        When a new player joins the game
        And there is no room waiting for another player
        Then that player should be put in a new room waiting for another player

    Scenario: Second player joins the game
        Given that there is a room waiting for another player with a player inside  
        When a new player joins the game
        Then that new player should be put in that room 
        And the room should change to waiting for the match to begin
        And the player inside the room should be notified about these changes

    Scenario: Third player joins the game
        Given that there is a room waiting for the match to begin with two players inside
        And there is no room waiting for another player
        When a new player joins the game
        Then the new player should be put in a new room and wait for another player

    Scenario: First player leaves the room
        Given that the player is in a room waiting for the match to begin with another player
        When one player decides to leave the room 
        Then the room should go back to waiting for another player
        And the player that remains in the room should be notified about these changes

    Scenario: Fourth player joins the game
        Given that there are two rooms waiting for another player, each with one player inside
        When a new player joins the game
        Then the new player should be put in the oldest room waiting for another player
        And the player that was already inside should be notified about these changes

    Scenario: Second player leaves the room
        Given that there is a room waiting for another player with a player inside  
        When that player decides to leave the room 
        Then the room should be destroied

