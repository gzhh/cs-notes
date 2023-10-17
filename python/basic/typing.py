from typing import List, Set, Dict, \
    Optional, Iterable


def get_match_list(match_ids: List[int]) -> List[Dict]:
    ret = list()
    for match_id in match_ids:
        item = dict()
        item['id'] = match_id
        item['name'] = '{}{}'.format('name', match_id)
        ret.append(item)
    return ret


def get_matches(match_ids: List[int]) -> Dict[int, Dict]:
    ret = dict()
    for match_id in match_ids:
        item = dict()
        item['id'] = match_id
        item['name'] = '{}{}'.format('name', match_id)
        ret[match_id] = item
    return ret


match_ids = [1, 2, 3, 4, 5]
# match_ids = list([1,2,3,4,5])
# match_ids = [i for i in range(5)]

match_list = get_match_list(match_ids)
print(match_list)

matches = get_matches(match_ids)
print(matches)
