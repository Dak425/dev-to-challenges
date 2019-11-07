class Word
  @@alphabet = ('a'..'z').to_a

  def self.to_mark(word)
    word.chars.sum { |char| @@alphabet.index(char.downcase) + 1 }
  end
end
