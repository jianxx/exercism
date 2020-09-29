def to_rna(dna_strand):
    transcribed = {"G": "C", "C": "G", "T": "A", "A": "U"}
    return ''.join([transcribed[n] for n in dna_strand])
