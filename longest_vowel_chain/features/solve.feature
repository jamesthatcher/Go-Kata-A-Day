Feature: Longest vowel chain
  The vowel substrings in the word codewarriors are o,e,a,io.
  The longest of these has a length of 2. Given a lowercase string that
  has alphabetic characters only and no spaces, return the length of the
  longest vowel substring.

  Scenario Outline: Solve
    Given I use the the <word>
    When I count the consecutive vowels
    Then it should return the <count> of the longest vowel chain

    Examples:
      | word                  | count |
      | codewarriors          | 2     |
      | suoidea               | 3     |
      | ultrarevolutionariees | 3     |
      | strengthlessnesses    | 1     |
      | cuboideonavicuare     | 2     |
      | chrononhotonthuooaos  | 5     |
      | iiihoovaeaaaoougjyaw  | 8     |