/*/
let p(w) wining probability. To calculate it we can use the following logistic function:
    l (x)  = 1 / 1 + 10 ^ [(score(a) - score(b)) / 400]

let s actual elo score (rating)
let k-factor be a "constant" value for the rating update speed of our players

(rating differential) Δr = (s - p(w)) * k-factor 
/*/

// evaluates logistic function in rating[elo(a)-elo(b)]:
// 1 / 10 ^ -(rating[A-B]/ 400) + 1
float win_probability(short player_rating_a, short player_rating_b);

// returns the elo differential for player a calibrated player A
short rating_differential(short player_rating_a, short player_rating_b, float key_factor);
