package words

var adjectives = []string{
	"able", "abnormal", "above average", "absent-minded", "accepting", "adaptable", "adventurous", "affable", "affectionate", "aggressive", "agile", "agreeable", "alert", "aloof", "amazing", "ambitious", "amiable", "amicable", "amusing", "analytical", "angelic", "annoying", "apathetic", "apprehensive", "ardent", "arrogant", "articulate", "artificial", "artistic", "assertive", "athletic", "attentive", "average", "awesome", "awful", "awkward", "balanced", "beautiful", "belligerent", "below average", "beneficent", "big-headed", "bitchy", "blue", "blunt", "boastful", "boisterous", "bold", "bone-idle", "boring", "bossy", "brave", "bright", "brilliant", "broad-minded", "buff", "busy", "callous", "calm", "candid", "cantankerous", "capable", "careful", "careless", "caring", "caustic", "cautious", "changeable", "charming", "cheerful", "chic", "childish", "childlike", "churlish", "circumspect", "civil", "clean", "clever", "clinging", "clumsy", "coherent", "cold", "communicative", "compassionate", "competent", "complex", "composed", "compulsive", "conceited", "condescending", "confident", "confused", "conscientious", "conservative", "considerate", "content", "convivial", "cool", "cool-headed", "cooperative", "cordial", "courageous", "courteous", "cowardly", "crabby", "crafty", "cranky", "crass", "creative", "critical", "cruel", "cunning", "curious", "cynical", "dainty", "daring", "dark", "deceitful", "decisive", "deep", "deferential", "defiant", "deft", "delicate", "delightful", "demanding", "demonic", "demure", "dependent", "depressed", "detached", "determined", "devoted", "devout", "dextrous", "diligent", "diplomatic", "direct", "dirty", "disagreeable", "discerning", "discreet", "disgruntled", "dishonest", "disruptive", "distant", "distraught", "distrustful", "dogmatic", "domineering", "dowdy", "dramatic", "dreamer", "dreary", "drowsy", "drugged", "drunk", "dull", "dutiful", "dynamic", "eager", "earnest", "easy-going", "easygoing", "efficient", "egotistical", "elfin", "embarrassed", "emotional", "energetic", "enterprising", "enthusiastic", "evasive", "even-tempered", "exacting", "excellent", "excitable", "excited", "experienced", "expert", "exuberant", "fabulous", "fair", "fair-minded", "faithful", "fancy", "fastidious", "fearless", "ferocious", "fervent", "fiery", "fighter", "finicky", "flabby", "flaky", "flashy", "flirtatious", "foolhardy", "foolish", "forceful", "forgiving", "frank", "free", "friendly", "frustrated", "fun-loving", "funny", "fussy", "generous", "gentle", "giving", "gloomy", "glutinous", "good", "gorgeous", "gracious", "grave", "great", "greedy", "gregarious", "groggy", "grouchy", "grumpy", "guarded", "gullible", "handsome", "happy", "hard-working", "harsh", "hateful", "hearty", "helpful", "hesitant", "honest", "hopeful", "hot-headed", "humble", "humorous", "hypercritical", "hysterical", "idiotic", "idle", "illogical", "imaginative", "immature", "immodest", "impartial", "impatient", "imperturbable", "impetuous", "impolite", "impractical", "impressionable", "impressive", "impulsive", "inactive", "incisive", "incompetent", "inconsiderate", "inconsistent", "indecisive", "indefatigable", "independent", "indiscreet", "indolent", "industrious", "inexperienced", "inflexible", "insensitive", "inspiring", "intellectual", "intelligent", "interesting", "interfering", "intolerant", "intuitive", "inventive", "irascible", "irresponsible", "irritable", "irritating", "jealous", "jocular", "jovial", "joyful", "joyous", "judgmental", "keen", "kind", "knowledgeable", "lame", "lazy", "leader", "lean", "leery", "lethargic", "level-headed", "light", "light-hearted", "likeable", "listless", "lithe", "lively", "local", "logical", "long-winded", "lovable", "love-lorn", "lovely", "loving", "loyal", "machiavellian", "manipulative", "materialistic", "maternal", "mature", "mean", "meddlesome", "melancholy", "mercurial", "merry", "messy", "methodical", "meticulous", "mild", "mischievous", "miserable", "miserly", "modest", "moody", "moronic", "morose", "motivated", "musical", "naive", "narrow-minded", "nasty", "natural", "naughty", "naïve", "neat", "negative", "nervous", "nice", "noisy", "normal", "nosy", "numb", "obliging", "obnoxious", "obsessive", "obstinate", "old-fashioned", "one-sided", "opinionated", "optimistic", "orderly", "organized", "ostentatious", "outgoing", "outspoken", "overcritical", "overemotional", "parsimonious", "passionate", "passive", "paternal", "paternalistic", "patient", "patriotic", "patronizing", "peaceful", "peevish", "pensive", "perfectionist", "persevering", "persistent", "persnickety", "personable", "perverse", "pessimistic", "petulant", "philosophical", "picky", "pioneering", "pitiful", "placid", "plain", "plain-speaking", "playful", "pleasant", "pleasing", "plucky", "polite", "pompous", "poor", "popular", "positive", "possessive", "powerful", "practical", "prejudiced", "pretty", "prim", "pro-active", "proficient", "proper", "proud", "provocative", "prudent", "punctual", "pusillanimous", "quarrelsome", "querulous", "questioning", "quick", "quick-tempered", "quick-witted", "quiet", "radical", "rational", "realistic", "reassuring", "rebellious", "reclusive", "reflective", "relaxed", "reliable", "religious", "reluctant", "resentful", "reserved", "resigned", "resourceful", "respected", "respectful", "responsible", "restless", "revered", "reverent", "rich", "ridiculous", "rigid", "romantic", "rude", "ruthless", "sad", "sarcastic", "sassy", "saucy", "secretive", "sedate", "self-assured", "self-centred", "self-confident", "self-conscious", "self-disciplined", "self-indulgent", "selfish", "sensible", "sensitive", "sentimental", "serene", "serious", "sharp", "short", "short-tempered", "shrewd", "shy", "silly", "simple", "simple-minded", "sincere", "sleepy", "slight", "sloppy", "slothful", "slovenly", "slow", "smart", "snazzy", "sneaky", "sneering", "snobby", "sober", "sociable", "somber", "sophisticated", "soulful", "soulless", "sour", "spirited", "spiteful", "stable", "staid", "steady", "stern", "stingy", "stoic", "straightforward", "striking", "strong", "stubborn", "studious", "stupid", "sturdy", "subtle", "successful", "sulky", "sullen", "supercilious", "superficial", "surly", "suspicious", "sweet", "sympathetic", "tactful", "tactless", "talented", "tall", "tantalizing", "tender", "tense", "testy", "thinking", "thoughtful", "thoughtless", "thrilling", "tidy", "timid", "tired", "tireless", "tolerant", "touchy", "tough", "tranquil", "tricky", "truculent", "trusting", "ugly", "unaffected", "unassuming", "unbalanced", "uncertain", "uncooperative", "undependable", "understanding", "unemotional", "unfriendly", "unguarded", "unhappy", "unhelpful", "unimaginative", "unique", "unkind", "unlucky", "unmotivated", "unpleasant", "unpopular", "unpredictable", "unreliable", "unselfish", "unsophisticated", "unstable", "unsure", "unthinking", "untidy", "untrustworthy", "unwilling", "vague", "vain", "venal", "vengeful", "versatile", "vigilant", "volcanic", "vulgar", "vulnerable", "warm", "warmhearted", "wary", "watchful", "weak", "weak-willed", "wild", "willing", "wise", "witty", "wonderful", "zealous",
}

var animals = []string{
	"albatross", "alligator", "angelfish", "ant", "anteater", "antelope", "armadillo", "avocet", "axolotl", "aye-aye", "baboon", "badger", "bandicoot", "barb", "barnacle", "barracuda", "bat", "beagle", "bear", "beaver", "bee", "beetle", "bird", "bison", "bobcat", "bonobo", "buffalo", "bulldog", "bumblebee", "butterfly", "camel", "capybara", "cat", "caterpillar", "catfish", "centipede", "chameleon", "chamois", "cheetah", "chicken", "chihuahua", "chimpanzee", "chinchilla", "chipmunk", "cichlid", "clam", "coati", "cockroach", "collie", "coral", "corgi", "cow", "coyote", "crab", "crocodile", "cuscus", "cuttlefish", "dachshund", "dalmatian", "deer", "dingo", "discus", "dodo", "dog", "dolphin", "donkey", "dormouse", "dragon", "dragonfly", "duck", "dugong", "eagle", "echidna", "eel", "elephant", "emu", "falcon", "ferret", "fish", "flamingo", "flounder", "fly", "fox", "frog", "gecko", "gibbon", "giraffe", "gnu", "goat", "goose", "gopher", "gorilla", "grasshopper", "grouse", "guppy", "hamster", "hare", "hedgehog", "hippopotamus", "hornet", "horse", "human", "hummingbird", "husky", "hyena", "ibis", "insect", "jackal", "jackalope", "jaguar", "jellyfish", "kangaroo", "kiwi", "koala", "labradoodle", "labrador", "ladybird", "lemming", "lemur", "leopard", "liger", "lion", "lionfish", "lizard", "llama", "lobster", "lynx", "macaque", "macaw", "magpie", "manatee", "mandrill", "markhor", "marmoset", "mastiff", "mayfly", "meerkat", "millipede", "mole", "mongoose", "monkey", "monster", "moorhen", "moose", "moth", "mouse", "mule", "neanderthal", "newt", "nightingale", "numbat", "ocelot", "octopus", "okapi", "olm", "opossum", "orang-utan", "ostrich", "otter", "owl", "oyster", "pademelon", "panda", "panther", "parrot", "peacock", "pelican", "penguin", "pheasant", "pig", "pika", "pike", "pinscher", "piranha", "platypus", "poodle", "porcupine", "prawn", "puffin", "pug", "puma", "quail", "quetzal", "quokka", "quoll", "rabbit", "raccoon", "ragdoll", "rat", "rattlesnake", "reindeer", "rhinoceros", "robin", "rottweiler", "salamander", "saola", "scorpion", "seahorse", "seal", "shark", "sheep", "shrimp", "skunk", "sloth", "slug", "snail", "snake", "sparrow", "spider", "sponge", "squid", "squirrel", "starfish", "stingray", "swan", "tamarin", "tapir", "tarantula", "tarsier", "termite", "terrier", "tiger", "toad", "tortoise", "toucan", "tropicbird", "tuatara", "turkey", "turtle", "unicorn", "vulture", "wallaby", "walrus", "warthog", "wasp", "weasel", "whale", "wolf", "wombat", "woodlouse", "woodpecker", "worm", "yak", "zebra",
}
