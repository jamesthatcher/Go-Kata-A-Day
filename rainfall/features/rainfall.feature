Feature: Rainfall
  Data and data1 are two strings with rainfall records of a few cities for months from January to December. The records of towns are separated by \n. The name of each town is followed by :.

  data and towns can be seen in "Your Test Cases:".

  Task:
  function: mean(town, string) should return the average of rainfall for the city town and the strng data or data1 (In R and Julia this function is called avg).
  function: variance(town, string) should return the variance of rainfall for the city town and the strng data or data1.

  Scenario Outline: Mean
    Given I have data of type <data> and city <city>
    When the mean is calculated
    Then I should have an average rainfall of <avg-rainfall>

    Examples:
      | data                                                                                     | city           | avg-rainfall |
      | "London:Jan 48.0,Feb 38.0,Mar 39.2,Apr 42.2 Beijing:Jan 48.0,Feb 38.9,Mar 39.9,Apr 42.2" | "London"       | 41.85        |
      | "Beijing:Jan 98.0,Feb 38.9,Mar 39.9,Apr 42.2 London:Jan 68.0,Feb 38.0,Mar 39.2,Apr 42.2" | "Beijing"      | 56.85        |
      | "London:Jan 48.0,Feb 38.0,Mar 39.2,Apr 42.2 Beijing:Jan 48.0,Feb 38.9,Mar 39.9,Apr 42.2" | "doesnotexist" | -1           |


  Scenario Outline: Variance
    Given I have data of type <data> and city <city>
    When the mean is calculated
    Then I should have an average rainfall of <var-rainfall>

    Examples:
      | data                                                                                     | city           | var-rainfall |
      | "London:Jan 48.0,Feb 38.0,Mar 39.2,Apr 42.2 Beijing:Jan 48.0,Feb 38.9,Mar 39.9,Apr 42.2" | "London"       | 14.9475      |
      | "Beijing:Jan 58.0,Feb 38.9,Mar 39.9,Apr 42.2 London:Jan 48.0,Feb 38.0,Mar 39.2,Apr 42.2" | "Beijing"      | 51.4475      |
      | "London:Jan 48.0,Feb 38.0,Mar 39.2,Apr 42.2 Beijing:Jan 48.0,Feb 38.9,Mar 39.9,Apr 42.2" | "doesnotexist" | -1           |
