select h.hacker_id, 
   h.name 
from   submissions s
   join challenges c 
     on s.challenge_id = c.challenge_id 
   join difficulty d 
     on c.difficulty_level = d.difficulty_level 
   join hackers h 
     on s.hacker_id = h.hacker_id 
        and s.score = d.score 
group  by h.hacker_id, 
  h.name 
having count(s.hacker_id) > 1 
order  by count(s.hacker_id) desc,  s.hacker_id asc;