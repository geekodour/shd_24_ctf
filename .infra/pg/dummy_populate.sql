-- psql -d shd -f ./dummy_populate.sql
INSERT INTO team (id, team_name, secret, map_url)
VALUES
    ('t1', 'adenine', 'atcg', 'https://www.google.com/maps'),
    ('t2', 'cytosine', 'gcat', 'https://www.google.com/maps'),
    ('t3', 'guanine', 'tacg', 'https://www.google.com/maps'),
    ('t4', 'thymine', 'agct', 'https://www.google.com/maps'),
    ('t5', 'zanthine', 'cgta', 'https://www.google.com/maps');

INSERT INTO flag (id, question, answer, answer_option, score, seq_num)
VALUES
    ('f1aatcg', 'Which organelle is responsible for cellular respiration?', 'Mitochondria', '["Mitochondria", "Golgi Body", "Endoplasmic Reticulum", "Lysosome"]'::jsonb, 100, 1),
    ('f2attttcg', 'What is the pH of a strong acid?', '0-3', '["0-3", "7", "11-14", "4-6"]'::jsonb, 150, 2),
    ('f3atccg', 'What is the time complexity of binary search?', 'O(log n)', '["O(n)", "O(log n)", "O(n²)", "O(1)"]'::jsonb, 100, 3),
    ('f4atcgg', 'According to F = ma, what happens to acceleration if force doubles but mass stays constant?', 'Acceleration doubles', '["Acceleration halves", "Acceleration doubles", "Acceleration stays the same", "Acceleration quadruples"]'::jsonb, 100, 4),
    ('f5aatcg', 'What happens to current when resistance increases in a circuit with constant voltage?', 'Current decreases', '["Current increases", "Current decreases", "Current stays the same", "Current becomes zero"]'::jsonb, 125, 5),

    ('f6gccat', 'Which process converts glucose to pyruvate?', 'Glycolysis', '["Glycolysis", "Krebs Cycle", "Beta Oxidation", "Gluconeogenesis"]'::jsonb, 100, 1),
    ('f7gcaat', 'What type of bond forms when electrons are shared equally?', 'Covalent', '["Ionic", "Covalent", "Metallic", "Hydrogen"]'::jsonb, 150, 2),
    ('f8gcaatt', 'What does HTML stand for?', 'Hypertext Markup Language', '["Hypertext Markup Language", "High Technical Modern Language", "Hypertext Management Logic", "High Text Marking Language"]'::jsonb, 100, 3),
    ('f9gggcat', 'What is the SI unit of pressure?', 'Pascal', '["Newton", "Pascal", "Joule", "Watt"]'::jsonb, 100, 4),
    ('f10gcatt', 'Which component amplifies electrical signals?', 'Transistor', '["Diode", "Transistor", "Resistor", "Capacitor"]'::jsonb, 125, 5),

    ('f11taacg', 'What is the primary function of RNA?', 'Protein synthesis', '["DNA replication", "Protein synthesis", "Energy storage", "Cell division"]'::jsonb, 100, 1),
    ('f12tacccg', 'Which gas makes solutions acidic when dissolved in water?', 'Carbon dioxide', '["Oxygen", "Nitrogen", "Carbon dioxide", "Hydrogen"]'::jsonb, 150, 2),
    ('f13tacgg', 'What is the purpose of a stack in programming?', 'Last In First Out data storage', '["Random data access", "Last In First Out data storage", "Permanent data storage", "Parallel processing"]'::jsonb, 100, 3),
    ('f14ttacccgg', 'What is the unit of electric potential difference?', 'Volt', '["Ampere", "Volt", "Ohm", "Watt"]'::jsonb, 100, 4),
    ('f15taacggg', 'What type of circuit has only one path for current?', 'Series', '["Parallel", "Series", "Complex", "Hybrid"]'::jsonb, 125, 5),

    ('f16aagct', 'What process produces ATP without oxygen?', 'Anaerobic respiration', '["Aerobic respiration", "Anaerobic respiration", "Photosynthesis", "Fermentation"]'::jsonb, 100, 1),
    ('f17aggcct', 'What is the charge of an electron?', 'Negative', '["Positive", "Negative", "Neutral", "Variable"]'::jsonb, 150, 2),
    ('f18agctt', 'Which data structure uses FIFO principle?', 'Queue', '["Stack", "Queue", "Array", "Tree"]'::jsonb, 100, 3),
    ('f19agct', 'What is the formula for kinetic energy?', '1/2 mv²', '["mgh", "1/2 mv²", "F=ma", "P=mv"]'::jsonb, 100, 4),
    ('f20aagctt', 'What does AC stand for in electronics?', 'Alternating Current', '["Alternating Current", "Active Current", "Ampere Current", "Actual Current"]'::jsonb, 125, 5),

    ('f21cggta', 'What is the function of chlorophyll?', 'Light absorption', '["Light absorption", "Sugar storage", "Water storage", "Gas exchange"]'::jsonb, 100, 1),
    ('f22cgtta', 'What happens in an exothermic reaction?', 'Heat is released', '["Heat is absorbed", "Heat is released", "No temperature change", "Mass is lost"]'::jsonb, 150, 2),
    ('f23ccccgta', 'What is the smallest unit of digital information?', 'Bit', '["Byte", "Bit", "Kilobyte", "Pixel"]'::jsonb, 100, 3),
    ('f24cgtaaa', 'What is the function of a convex lens?', 'Converges light', '["Diverges light", "Converges light", "Blocks light", "Splits light"]'::jsonb, 100, 4),
    ('f25cgta', 'What component opposes change in current?', 'Inductor', '["Capacitor", "Inductor", "Resistor", "Diode"]'::jsonb, 125, 5);
