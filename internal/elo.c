#ifndef ELO_H_INCLUDED
#define ELO_H_INCLUDED
#include "elo.h"
#include <math.h>

float win_probability(short player_rating_a, short player_rating_b)
{
    return 1 / (pow(10, ((double)-(player_rating_a - player_rating_b) / 400)) + 1);
}

short rating_differential(short player_rating_a, short player_rating_b, float population_growth_factor)
{
    return round((player_rating_a - win_probability(player_rating_a, player_rating_b) * population_growth_factor));
}

#endif