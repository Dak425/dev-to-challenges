require 'minitest/autorun'
require_relative 'word'

class TestWords < Minitest::Test
  def setup
    @cases = [
      { input: 'love', expected: 54 },
      { input: 'friendship', expected: 108 },
      { input: 'attitude', expected: 100 },
      { input: 'friends', expected: 75 },
      { input: 'family', expected: 66 },
      { input: 'selflessness', expected: 154 },
      { input: 'knowledge', expected: 96 }
    ]
  end

  def test_to_mark
    @cases.each { |test| assert_equal test[:expected], Word.to_mark(test[:input]) }
  end
end
