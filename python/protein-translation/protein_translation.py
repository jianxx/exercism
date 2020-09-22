def proteins(strand):
    translater = {
        "AUG": "Methionine",
        "UUU": "Phenylalanine",
        "UUC": "Phenylalanine",
        "UUA": "Leucine",
        "UUG": "Leucine",
        "UCU": "Serine",
        "UCC": "Serine",
        "UCA": "Serine",
        "UCG": "Serine",
        "UAU": "Tyrosine",
        "UAC": "Tyrosine",
        "UGU": "Cysteine",
        "UGC": "Cysteine",
        "UGG": "Tryptophan",
        "UAA": "STOP",
        "UAG": "STOP",
        "UGA": "STOP",
    }
    proteins = []
    for codon in [strand[i : i + 3] for i in range(0, len(strand), 3)]:
        if translater.get(codon, None):
            if translater.get(codon, None) == "STOP":
                break
            proteins.append(translater.get(codon, None))
    return proteins
