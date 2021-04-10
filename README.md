# ldapMon

Gets contextCSN from two given ldap servers

## Config

cred.json stores configuration edit it to fit your needs

## Test

You can test it by running 3 ldap containers

```
# Create the first ldap server, save the container id in LDAP_CID and get its IP:

docker run -d --name ldap_master1 --rm -p 4000:389  --env LDAP_ORGANISATION="organization" --env LDAP_DOMAIN="organization.com" --env LDAP_ADMIN_PASSWORD="JhonSnow" --env LDAP_BASE_DN="dc=organization,dc=com" --env LDAP_REPLICATION=true --hostname ldap.example.org osixia/openldap:1.2.4```

#Second ldap

docker run -d --name ldap_master2 --rm -p 4002:389  --env LDAP_ORGANISATION="organization" --env LDAP_DOMAIN="organization.com" --env LDAP_ADMIN_PASSWORD="JhonSnow" --env LDAP_BASE_DN="dc=organization,dc=com" --env LDAP_REPLICATION=true --hostname ldap2.example.org osixia/openldap:1.2.4

# Set resolution between servers

docker exec ldap_master1 bash -c "echo 172.17.0.3 ldap2.example.org >> /etc/hosts"
docker exec ldap_master2 bash -c "echo 172.17.0.2 ldap.example.org >> /etc/hosts"

If you have more containers ensure what ips got ldaps with docker inspect
```
