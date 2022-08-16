import yaml

with open("configs.yaml", 'r') as conf_file:
    config  = yaml.load(conf_file, Loader=yaml.FullLoader)

print(config)