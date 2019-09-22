Feature: Translate a string to Pilot's alphabet (NATO phonetic alphabet)
  The idea for this Kata came from 9gag today.here

  You'll have to translate a string to Pilot's alphabet (NATO phonetic alphabet) wiki.

  Like this:

  Input: If you can read

  Output: India Foxtrot Yankee Oscar Uniform Charlie Alfa November Romeo Echo Alfa Delta

  Keep the punctuation, and remove the spaces.
  Use Xray without dash or space.

  Scenario Outline: Tests using hard-coded strings
    Given a word <str>
    Then it should return a correct <phonetic> translation of the <str>

    Examples:
      | str                       | phonetic                                                                                                                |
      | "If you can read"         | "India Foxtrot Yankee Oscar Uniform Charlie Alfa November Romeo Echo Alfa Delta"                                        |
      | "Did not see that coming" | "Delta India Delta November Oscar Tango Sierra Echo Echo Tango Hotel Alfa Tango Charlie Oscar Mike India November Golf" |
      | "go for it!"              | "Golf Oscar Foxtrot Oscar Romeo India Tango !"                                                                          |