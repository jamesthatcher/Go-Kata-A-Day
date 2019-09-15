Feature: Generator
  My friend wants a new band name for her band. She like bands that use the formula:
  "The" + a noun with the first letter capitalized, for example:

  "dolphin" -> "The Dolphin"

  However, when a noun STARTS and ENDS with the same letter, she likes to repeat the
  noun twice and connect them together with the first and last letter, combined into
  one word (WITHOUT "The" in front), like this:

  "alaska" -> "Alaskalaska"

  Complete the function that takes a noun as a string, and returns her preferred band
  name written as a string.

  Scenario Outline: name noun without START and END matching
    Given I have <names> without the start and end matching
    When I generate
    Then it should return <names> of the <band>

    Examples:
      | names     | band          |
      | "knife"   | "The Knife"   |
      | "bed"     | "The Bed"     |
      | "Cheese"  | "The Cheese"  |
      | "pencil"  | "The Pencil"  |
      | "onion"   | "The Onion"   |
      | "Cookie"  | "The Cookie"  |
      | "dolphin" | "The Dolphin" |

  Scenario Outline: name noun with START and END matching
    Given I have <names> with the start and end matching
    When I generate
    Then it should return <names> of the <band>

    Examples:
      | names     | band            |
      | "tart"    | "Tartart"       |
      | "sandles" | "Sandlesandles" |
      | "ss"      | "Sss"          |
      | "snakes"  | "Snakesnakes"   |
      | "seconds" | "Secondseconds" |
      | "tght"    | "Tghtght"       |
      | "alaska"  | "Alaskalaska"   |