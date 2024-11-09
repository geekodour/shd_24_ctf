-- psql -d shd -f ./dummy_populate.sql
INSERT INTO team (id, team_name, secret, map_url)
VALUES
    ('t1', 'CyberPhoenix', 'ph03n1x_r1s3s_2024', ''),
    ('t2', 'ByteBusters', 'bust3r5_un1t3_9876', ''),
    ('t3', 'HackMasters', 'mast3r_0f_h4cks_42', ''),
    ('t4', 'NetRunners', 'runn3r5_3dg3_1337', ''),
    ('t5', 'CodeSentinels', 's3nt1n3l_gu4rd_789', '');

INSERT INTO flag (id, question, answer, answer_option, score, seq_num)
VALUES
    ('f1', 'What is the most common port number for HTTPS?', '443', '["443", "80", "8080", "8443"]'::jsonb, 100, 1),
    ('f2', 'Which encryption algorithm is used in WPA3?', 'SAE (Simultaneous Authentication of Equals)', '["WEP", "TKIP", "SAE (Simultaneous Authentication of Equals)", "RC4"]'::jsonb, 150, 2),
    ('f3', 'What command is used to list network interfaces in Linux?', 'ifconfig', '["ipconfig", "ifconfig", "netstat", "networkctl"]'::jsonb, 100, 3),
    ('f4', 'What is the default port for SSH?', '22', '["21", "22", "23", "25"]'::jsonb, 100, 4),
    ('f5', 'Which protocol is connectionless?', 'UDP', '["TCP", "UDP", "HTTP", "FTP"]'::jsonb, 125, 5),

    ('f6', 'What year was the first computer virus created?', '1971', '["1971", "1981", "1991", "2001"]'::jsonb, 200, 1),
    ('f7', 'What is the maximum length of an IPv4 address in bits?', '32', '["16", "32", "64", "128"]'::jsonb, 100, 2),
    ('f8', 'Which hash function is considered cryptographically broken?', 'MD5', '["SHA-256", "MD5", "SHA-3", "BLAKE2"]'::jsonb, 150, 3),
    ('f9', 'What is the default subnet mask for a Class C network?', '255.255.255.0', '["255.0.0.0", "255.255.0.0", "255.255.255.0", "255.255.255.255"]'::jsonb, 125, 4),
    ('f10', 'Which port is commonly used for DNS?', '53', '["53", "25", "110", "143"]'::jsonb, 100, 5),

    ('f11', 'What does CSRF stand for?', 'Cross-Site Request Forgery', '["Cross-Site Request Forgery", "Common System Response Format", "Client Server Request Format", "Cross-Site Reference Frame"]'::jsonb, 175, 1),
    ('f12', 'Which protocol operates at the Network layer of the OSI model?', 'IP', '["TCP", "IP", "HTTP", "FTP"]'::jsonb, 150, 2),
    ('f13', 'What is the default password for Kali Linux?', 'kali', '["root", "toor", "kali", "admin"]'::jsonb, 100, 3),
    ('f14', 'Which command shows running processes in Linux?', 'ps', '["ps", "ls", "top", "proc"]'::jsonb, 100, 4),
    ('f15', 'What is the primary function of a firewall?', 'Network Security', '["Network Security", "Data Encryption", "User Authentication", "Data Compression"]'::jsonb, 125, 5),

    ('f16', 'Which port is used for HTTP?', '80', '["80", "443", "22", "21"]'::jsonb, 100, 1),
    ('f17', 'What does XSS stand for?', 'Cross-Site Scripting', '["Cross-Site Scripting", "Extended Security System", "XML Security Service", "X Server System"]'::jsonb, 150, 2),
    ('f18', 'Which protocol is used for secure email transmission?', 'SMTPS', '["SMTP", "SMTPS", "POP3", "IMAP"]'::jsonb, 125, 3),
    ('f19', 'What is the maximum theoretical bandwidth of CAT6 cable?', '10 Gbps', '["1 Gbps", "10 Gbps", "40 Gbps", "100 Gbps"]'::jsonb, 150, 4),
    ('f20', 'Which command is used to change file permissions in Linux?', 'chmod', '["chown", "chmod", "chgrp", "chattr"]'::jsonb, 100, 5),

    ('f21', 'What is the default gateway IP in most home networks?', '192.168.1.1', '["192.168.1.1", "10.0.0.1", "172.16.0.1", "127.0.0.1"]'::jsonb, 100, 1),
    ('f22', 'Which encryption algorithm is used in RSA?', 'Public Key', '["Symmetric Key", "Public Key", "Block Cipher", "Stream Cipher"]'::jsonb, 175, 2),
    ('f23', 'What does SQL injection attack target?', 'Database', '["Network", "Database", "Memory", "CPU"]'::jsonb, 150, 3),
    ('f24', 'Which protocol is used for secure web browsing?', 'HTTPS', '["HTTP", "HTTPS", "FTP", "SMTP"]'::jsonb, 100, 4),
    ('f25', 'What is the default port for FTP?', '21', '["20", "21", "22", "23"]'::jsonb, 100, 5);
