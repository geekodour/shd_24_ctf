-- psql -d shd -f ./dummy_populate.sql
INSERT INTO team (team_name, secret)
VALUES
    ('CyberPhoenix', 'ph03n1x_r1s3s_2024'),
    ('ByteBusters', 'bust3r5_un1t3_9876'),
    ('HackMasters', 'mast3r_0f_h4cks_42'),
    ('NetRunners', 'runn3r5_3dg3_1337'),
    ('CodeSentinels', 's3nt1n3l_gu4rd_789');

INSERT INTO flag (question, answer, answer_option, score, seq_num)
VALUES
    ('What is the most common port number for HTTPS?', '443', '["443", "80", "8080", "8443"]'::jsonb, 100, 1),
    ('Which encryption algorithm is used in WPA3?', 'SAE (Simultaneous Authentication of Equals)', '["WEP", "TKIP", "SAE (Simultaneous Authentication of Equals)", "RC4"]'::jsonb, 150, 2),
    ('What command is used to list network interfaces in Linux?', 'ifconfig', '["ipconfig", "ifconfig", "netstat", "networkctl"]'::jsonb, 100, 3),
    ('What is the default port for SSH?', '22', '["21", "22", "23", "25"]'::jsonb, 100, 4),
    ('Which protocol is connectionless?', 'UDP', '["TCP", "UDP", "HTTP", "FTP"]'::jsonb, 125, 5),
    ('What year was the first computer virus created?', '1971', '["1971", "1981", "1991", "2001"]'::jsonb, 200, 6),
    ('What is the maximum length of an IPv4 address in bits?', '32', '["16", "32", "64", "128"]'::jsonb, 100, 7),
    ('Which hash function is considered cryptographically broken?', 'MD5', '["SHA-256", "MD5", "SHA-3", "BLAKE2"]'::jsonb, 150, 8),
    ('What is the default subnet mask for a Class C network?', '255.255.255.0', '["255.0.0.0", "255.255.0.0", "255.255.255.0", "255.255.255.255"]'::jsonb, 125, 9),
    ('Which port is commonly used for DNS?', '53', '["53", "25", "110", "143"]'::jsonb, 100, 10),
    ('What does CSRF stand for?', 'Cross-Site Request Forgery', '["Cross-Site Request Forgery", "Common System Response Format", "Client Server Request Format", "Cross-Site Reference Frame"]'::jsonb, 175, 11),
    ('Which protocol operates at the Network layer of the OSI model?', 'IP', '["TCP", "IP", "HTTP", "FTP"]'::jsonb, 150, 12),
    ('What is the default password for Kali Linux?', 'kali', '["root", "toor", "kali", "admin"]'::jsonb, 100, 13),
    ('Which command shows running processes in Linux?', 'ps', '["ps", "ls", "top", "proc"]'::jsonb, 100, 14),
    ('What is the primary function of a firewall?', 'Network Security', '["Network Security", "Data Encryption", "User Authentication", "Data Compression"]'::jsonb, 125, 15),
    ('Which port is used for HTTP?', '80', '["80", "443", "22", "21"]'::jsonb, 100, 16),
    ('What does XSS stand for?', 'Cross-Site Scripting', '["Cross-Site Scripting", "Extended Security System", "XML Security Service", "X Server System"]'::jsonb, 150, 17),
    ('Which protocol is used for secure email transmission?', 'SMTPS', '["SMTP", "SMTPS", "POP3", "IMAP"]'::jsonb, 125, 18),
    ('What is the maximum theoretical bandwidth of CAT6 cable?', '10 Gbps', '["1 Gbps", "10 Gbps", "40 Gbps", "100 Gbps"]'::jsonb, 150, 19),
    ('Which command is used to change file permissions in Linux?', 'chmod', '["chown", "chmod", "chgrp", "chattr"]'::jsonb, 100, 20),
    ('What is the default gateway IP in most home networks?', '192.168.1.1', '["192.168.1.1", "10.0.0.1", "172.16.0.1", "127.0.0.1"]'::jsonb, 100, 21),
    ('Which encryption algorithm is used in RSA?', 'Public Key', '["Symmetric Key", "Public Key", "Block Cipher", "Stream Cipher"]'::jsonb, 175, 22),
    ('What does SQL injection attack target?', 'Database', '["Network", "Database", "Memory", "CPU"]'::jsonb, 150, 23),
    ('Which protocol is used for secure web browsing?', 'HTTPS', '["HTTP", "HTTPS", "FTP", "SMTP"]'::jsonb, 100, 24),
    ('What is the default port for FTP?', '21', '["20", "21", "22", "23"]'::jsonb, 100, 25);
