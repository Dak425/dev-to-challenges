class Student
  attr_reader :fives, :tens, :twenties

  def initialize(fives, tens, twenties)
    @fives = fives
    @tens = tens
    @twenties = twenties
  end

  def funds
    (@fives * 5) + (@tens * 10) + (@twenties * 20)
  end
end
