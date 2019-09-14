Feature: Longest vowel chain
  The vowel substrings in the word codewarriors are o,e,a,io.
  The longest of these has a length of 2. Given a lowercase string that
  has alphabetic characters only and no spaces, return the length of the
  longest vowel substring.

  Scenario Outline: Solve
    Given I need to count the consecutive vowels
    When I count <words>
    Then it should return <count> of longest vowel chain in <words>

    Examples:
      | words                   | count |
      | "codewarriors"          | 2     |
      | "suoidea"               | 3     |
      | "ultrarevolutionariees" | 3     |
      | "strengthlessnesses"    | 1     |
      | "cuboideonavicuare"     | 2     |
      | "chrononhotonthuooaos"  | 5     |
      | "iiihoovaeaaaoougjyaw"  | 8     |