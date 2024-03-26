curl -s https://zero.academie.one/assets/superhero/all.json | jq -r '.[] | select(.id == 170) | .name'
curl -s https://zero.academie.one/assets/superhero/all.json | jq -r '.[] | select(.id == 170) | .powerstats.power'
curl -s https://zero.academie.one/assets/superhero/all.json | jq -r '.[] | select(.id == 170) | .appearance.gender'