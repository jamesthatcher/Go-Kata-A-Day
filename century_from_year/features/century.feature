Feature: Century
  The first century spans from the year 1 up to and including the year 100,
  The second - from the year 101 up to and including the year 200, etc.
  With a given year, return the century it is in.

  Scenario Outline: Century
    Given a year
    Then it should return <century> of the <year>

    Examples:
      | year                   | century |
      | 1990                   | 20      |
      | 1705                   | 18      |
      | 1900                   | 19      |
      | 1601                   | 17      |
      | 2000                   | 20      |
      | 89                     | 1       |