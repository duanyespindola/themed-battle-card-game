# Real Epic Card Battle

> by Duany Espindola

## Summary

This is a online, turn-based, combat card game, where the objective is to reduce your opponent's life points to zero.

## Understanding the Cards

### Common Attributes

The following attributes apply to all cards in the game:

#### Type

There are three types of cards: Attack Cards, Defense Cards, and Magic Cards.

Each type has a specific function and can only be used at specific times (see the [Turns](#Turns) chapter).

#### Cost

The value that will be immediately deducted from the player's Life Points when the card is played.

#### Name

The unique identification of the card.
Example: "Holly Lance," "Silver Shield," "Love Potion," etc.

#### Description

A brief text describing the card and its possible effects.

Example:
"The attack of this lance cannot be defended by cards with a cost lower than 5."

"This potion restores 10 points of the player's life and 5 points of the opponent's life."

#### Image

<figure>
      <img src="./doc%5Cimg/card-example01.png" width="100px"/>
      <figcaption>
            Illustrative image to assist in the customization/identification of the cards
      </figcaption>
</figure>

#### Attack Cards

Attack Cards have an attribute called **Attack Points**, which is the amount of points expected to decrease the opponent's _Life Points_ if no _Defense Move_ is performed and there are no active _Magic Effects_ influencing it.

Example:
"The 'Holly Lance' card with 10 attack points is played. The opponent does not perform a Defense Move, so their Life Points decrease by 10."

#### Defense Cards

Defense Cards have an attribute called **Defense Points**, which is the maximum amount of points expected to decrease from an opponent's _Attack Move_ if there are no active _Magic Effects_ influencing it.

Example:
"The opponent plays the 'Holly Lance' card with 10 attack points in their Attack Move. You play the defense card 'Silver Shield' with 9 Defense Points during your Defense Move, so your Life Points decrease by 1."

#### Magic Cards

Magic Cards do not have any additional attributes, but their description indicates the effect of the played card and when that effect will occur.

Example:
"The 'Love Potion' card prevents your opponent from performing an Attack Move in the next turn."

**Important:** The cost of the card is always "charged" at the moment it is played, and the effect is verified afterward. Therefore, for example, in cards that restore energy, care must be taken not to reduce one's own points to zero before receiving the recovery benefit.

Example:
"The 'Miraculous Healing' card costs 3 points and restores 10 Life Points. If the player has only 3 LP and uses this card, their LP will be reduced to zero first (resulting in a loss) before gaining the 10 LP from the card."

#### Number of Cards

##### Number of Cards in the Deck

Players must build a Deck (a set of cards) with 60 cards.

There is no minimum or maximum limit for the types of cards, and it is up to the player to balance the quantity of each type according to their game strategy.

##### Number of Cards in Hand

Before the start of the first turn, each player begins with 5 cards in hand. This is also the maximum number of cards a player can have in hand at the end of each Move within the turn.

At the end of each turn, each player draws a single card from their deck, unless the effect of a magic card alters this quantity.

Example:
"The player has 5 cards and draws one at the beginning of the turn. It's their
turn to make an Attack Move, but they don't have an Attack Card, so they must discard a card from their hand."

Example 2:
"The player has 5 cards in hand and plays a magic card (now having 4 cards in hand) that allows them to draw 3 cards (now having 7 cards in hand). At the end of the Magic Move, they must discard 2 cards."

##### Number of Active Magic Cards

Some Magic Cards may have an effect that lasts more than one turn, for example:

"This poisonous potion reduces your opponent's LP by 1 at the beginning of each attack action."

#### Turns

A turn consists of 3 different sequential moves:

1. A player starts the turn by making an Attack Move.
2. The opponent can then make a Defense Move.
3. After the Defense Move is completed and the damage is calculated, the attacking player can Cast a Magic.

Players switch roles each turn, meaning they act as the Attacking Player in one turn and as the Defending Player in the next.

### Elements

- **Stock** - The total number of cards a player possesses. A player can have an unlimited number of cards, whether repeated or not. The player needs to choose 60 cards from their Stock to form a Deck.
- **Deck** - The set of 60 cards chosen to start a game.
- **Hand** - The set of cards available to the player to use during the turn. Maximum of 5 cards.
- **Discard Pile** - The set of cards used or discarded during the game, which cannot be reused (unless the effect of a magic card states otherwise). Each player has their own discard pile containing only the cards from their own Deck.
- **Table** - The place where played cards are revealed.

# Want to contribute ?

[Guide for commits](./doc/git-templates.MD)
