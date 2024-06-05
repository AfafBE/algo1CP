// Algorithm to analyze a sentence character by character

// Initialisation des variables
sentence_length = 0
word_count = 1 (en supposant que la phrase commence par un mot)
vowel_count = 0

// Lire la phrase caractère par caractère
FOR EACH character IN sentence
  sentence_length++
// Vérifier si on a atteint la fin d'un mot
  IF character IS ' ' OR character IS '.' THEN
    word_count++ (indicates end of a word)
  END IF

  // Vérifier si c'est une voyelle
  IF character (converted to lowercase) IS 'a' OR character (converted to lowercase) IS 'e' OR 
     character (converted to lowercase) IS 'i' OR character (converted to lowercase) IS 'o' OR 
     character (converted to lowercase) IS 'u' THEN
    vowel_count++
  END IF
END FOR

// Afficher les résultats
PRINT "Nombre de caractères :", sentence_length
PRINT "Nombre de mots :", word_count
PRINT "Nombre de voyelles :", vowel_count
