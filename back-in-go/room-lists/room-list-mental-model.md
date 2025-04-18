Below, a simplified mental model of what I intend to implement, represented by a js object.
```javascript
{
    byStatus: {
        waitingAnotherPlayer : {
            "0f83398b-0599-40f6-b82d-65b218e8560f" : {
                id: "0f83398b-0599-40f6-b82d-65b218e8560f",
                status : "waitingAnotherPlayer",
                players : [p1, p2]
            },
            "05990f83-0599-40f6-b82d-65b218e8b82d" : {
                id: "05990f83-0599-40f6-b82d-65b218e8b82d",
                status : "waitingAnotherPlayer",
                players : [p1, p2]
            }
        }
        waitingMatchToStart : {
            "0f83398b-0599-40f6-b82d-65b218e8560f" : {
                id: "0f83398b-0599-40f6-b82d-65b218e8560f",
                status : "waitingMatchToStart",
                players : [p1, p2]
            },
            "05990f83-0599-40f6-b82d-65b218e8b82d" : {
                id: "05990f83-0599-40f6-b82d-65b218e8b82d",
                status : "waitingMatchToStart",
                players : [p1, p2]
            }
        }
    },
    byRoomId: {
        {
            "0f83398b-0599-40f6-b82d-65b218e8560f" : {
                id: "0f83398b-0599-40f6-b82d-65b218e8560f",
                status : "waitingAnotherPlayer",
                players : [p1, p2]
            },
            "05990f83-0599-40f6-b82d-65b218e8b82d" : {
                id: "05990f83-0599-40f6-b82d-65b218e8b82d",
                status : "waitingMatchToStart",
                players : [p1, p2]
            }            
        }
    }
    byPlayerId: {
            "05990f83-0599-40f6-b82d-65b218e8b82d" : {
                id: "0f83398b-0599-40f6-b82d-65b218e8560f",
                status : "waitingMatchToStart",
                players : [p1, p2]
            }
    }
    byConnection: {
            conn1 : {
                id: "0f83398b-0599-40f6-b82d-65b218e8560f",
                status : "waitingMatchToStart",'
                players : [p1, p2]
            },
            conn2 : {
                id: "0f83398b-0599-40f6-b82d-65b218e8560f",
                status : "waitingMatchToStart",
                players : [p1, p2]
            }
    }
}
```