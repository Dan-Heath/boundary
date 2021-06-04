begin;

  create table credential_vault_store (
    public_id wt_public_id primary key,
    scope_id wt_scope_id not null
      constraint iam_scope_fkey
        references iam_scope (public_id)
        on delete cascade
        on update cascade,
    name wt_name,
    description wt_description,
    create_time wt_timestamp,
    update_time wt_timestamp,
    version wt_version,
    vault_address wt_url not null,
    -- the remaining text columns can be null but if they are not null, they
    -- cannot contain an empty string
    namespace text
      constraint namespace_must_not_be_empty
        check(length(trim(namespace)) > 0),
    ca_cert bytea -- PEM encoded certificate bundle
      constraint ca_cert_must_not_be_empty
        check(length(ca_cert) > 0),
    tls_server_name text
      constraint tls_server_name_must_not_be_empty
        check(length(trim(tls_server_name)) > 0),
    tls_skip_verify boolean default false not null,
    constraint credential_store_fkey
      foreign key (scope_id, public_id)
      references credential_store (scope_id, public_id)
      on delete cascade
      on update cascade,
    constraint credential_vault_store_scope_id_name_uq
      unique(scope_id, name)
  );
  comment on table credential_vault_store is
    'credential_vault_store is a table where each row is a resource that represents a vault credential store. '
    'It is a credential_store subtype.';

  create trigger update_version_column after update on credential_vault_store
    for each row execute procedure update_version_column();

  create trigger update_time_column before update on credential_vault_store
    for each row execute procedure update_time_column();

  create trigger default_create_time_column before insert on credential_vault_store
    for each row execute procedure default_create_time();

  create trigger immutable_columns before update on credential_vault_store
    for each row execute procedure immutable_columns('public_id', 'scope_id','create_time');

  create trigger insert_credential_store_subtype before insert on credential_vault_store
    for each row execute procedure insert_credential_store_subtype();

  create trigger delete_credential_store_subtype after delete on credential_vault_store
    for each row execute procedure delete_credential_store_subtype();

  create table credential_vault_token_status_enm (
    name text primary key
      constraint only_predefined_token_statuses_allowed
      check (
        name in (
          'current',
          'maintaining',
          'revoked',
          'expired'
        )
      )
  );
  comment on table credential_vault_token_status_enm is
    'credential_vault_token_status_enm is an enumeration table for the status of vault tokens. '
    'It contains rows for representing the current, maintaining, revoked, and expired statuses.';

  insert into credential_vault_token_status_enm (name)
  values
    ('current'),
    ('maintaining'),
    ('revoked'),
    ('expired');

  create table credential_vault_token (
    token_hmac bytea primary key, -- hmac-sha256(token, key(blake2b-256(token_accessor))
    token bytea not null, -- encrypted value
    store_id wt_public_id not null
      constraint credential_vault_store_fkey
        references credential_vault_store (public_id)
        on delete cascade
        on update cascade,
    create_time wt_timestamp,
    update_time wt_timestamp,
    last_renewal_time timestamp with time zone not null,
    expiration_time timestamp with time zone not null
      constraint last_renewal_time_must_be_before_expiration_time
        check(last_renewal_time < expiration_time),
    key_id text not null
      constraint kms_database_key_version_fkey
        references kms_database_key_version (private_id)
        on delete restrict
        on update cascade,
    status text not null
      constraint credential_vault_token_status_enm_fkey
        references credential_vault_token_status_enm (name)
        on delete restrict
        on update cascade
  );
  comment on table credential_vault_token is
    'credential_vault_token is a table where each row contains a Vault token for one Vault credential store. '
    'A credential store can have only one vault token with the status of current';
  comment on column credential_vault_token.token_hmac is
    'token_hmac contains the hmac-sha256 value of the token. '
    'The hmac key is the blake2b-256 value of the token accessor.';

  -- https://www.postgresql.org/docs/current/indexes-partial.html
  create unique index credential_vault_token_current_status_constraint
    on credential_vault_token (store_id)
    where status = 'current';

  create index credential_vault_token_expiration_time_ix
    on credential_vault_token(expiration_time);
  comment on index credential_vault_token_expiration_time_ix is
    'the credential_vault_token_expiration_time_ix is used by the token renewal job';

  create trigger update_time_column before update on credential_vault_token
    for each row execute procedure update_time_column();

  create trigger default_create_time_column before insert on credential_vault_token
    for each row execute procedure default_create_time();

  create trigger immutable_columns before update on credential_vault_token
    for each row execute procedure immutable_columns('token_hmac', 'token', 'store_id','create_time');

  -- insert_credential_vault_token() is a before insert trigger
  -- function for credential_vault_token that changes the status of the current
  -- token to 'maintaining'
  create or replace function insert_credential_vault_token()
    returns trigger
  as $$
  begin
    update credential_vault_token
       set status   = 'maintaining'
     where store_id = new.store_id
       and status   = 'current';
    return new;
  end;
  $$ language plpgsql;

  create trigger insert_credential_vault_token before insert on credential_vault_token
    for each row execute procedure insert_credential_vault_token();

  create table credential_vault_client_certificate (
    store_id wt_public_id primary key
      constraint credential_vault_store_fkey
        references credential_vault_store (public_id)
        on delete cascade
        on update cascade,
    certificate bytea not null -- PEM encoded certificate
      constraint certificate_must_not_be_empty
        check(length(certificate) > 0),
    certificate_key bytea not null -- encrypted PEM encoded private key for certificate
      constraint certificate_key_must_not_be_empty
        check(length(certificate_key) > 0),
    certificate_key_hmac bytea not null
        constraint certificate_key_hmac_must_not_be_empty
            check(length(certificate_key_hmac) > 0),
    key_id text not null
      constraint kms_database_key_version_fkey
        references kms_database_key_version (private_id)
        on delete restrict
        on update cascade
  );
  comment on table credential_vault_client_certificate is
    'credential_vault_client_certificate is a table where each row contains a client certificate that a credential_vault_store uses for mTLS when connecting to Vault. '
    'A credential_vault_store can have 0 or 1 client certificates.';

  create trigger immutable_columns before update on credential_vault_client_certificate
    for each row execute procedure immutable_columns('store_id');

  create table credential_vault_http_method_enm (
    name text primary key
      constraint only_predefined_http_methods_allowed
      check (
        name in (
          'GET',
          'POST'
        )
      )
  );
  comment on table credential_vault_http_method_enm is
    'credential_vault_http_method_enm is an enumeration table for the http method used by a vault library when communicating with vault. '
    'It contains rows for representing the HTTP GET and the HTTP POST methods.';

  insert into credential_vault_http_method_enm (name)
  values
    ('GET'),
    ('POST');

  create table credential_vault_library (
    public_id wt_public_id primary key,
    store_id wt_public_id not null
      constraint credential_vault_store_fkey
        references credential_vault_store (public_id)
        on delete cascade
        on update cascade,
    name wt_name,
    description wt_description,
    create_time wt_timestamp,
    update_time wt_timestamp,
    version wt_version,
    vault_path text not null
      constraint vault_path_must_not_be_empty
        check(length(trim(vault_path)) > 0),
    http_method text not null
      constraint credential_vault_http_method_enm_fkey
        references credential_vault_http_method_enm (name)
        on delete restrict
        on update cascade,
    http_request_body bytea
      constraint http_request_body_only_allowed_with_post_method
        check(
          http_request_body is null
          or
          (
            http_method = 'POST'
            and
            length(http_request_body) > 0
          )
        ),
    constraint credential_vault_library_store_id_name_uq
      unique(store_id, name),
    constraint credential_library_fkey
      foreign key (store_id, public_id)
      references credential_library (store_id, public_id)
      on delete cascade
      on update cascade,
    constraint credential_vault_library_store_id_public_id_uq
      unique(store_id, public_id)
  );
  comment on table credential_vault_library is
    'credential_vault_library is a table where each row is a resource that represents a vault credential library. '
    'It is a credential_library subtype and a child table of credential_vault_store.';

  create trigger update_version_column after update on credential_vault_library
    for each row execute procedure update_version_column();

  create trigger update_time_column before update on credential_vault_library
    for each row execute procedure update_time_column();

  create trigger default_create_time_column before insert on credential_vault_library
    for each row execute procedure default_create_time();

  create trigger immutable_columns before update on credential_vault_library
    for each row execute procedure immutable_columns('public_id', 'store_id','create_time');

  create trigger insert_credential_library_subtype before insert on credential_vault_library
    for each row execute procedure insert_credential_library_subtype();

  create trigger delete_credential_library_subtype after delete on credential_vault_library
    for each row execute procedure delete_credential_library_subtype();

  create table credential_vault_credential_status_enm (
    name text primary key
      constraint only_predefined_credential_statuses_allowed
      check (
        name in (
          'active',
          'revoke',
          'revoked',
          'expired',
          'unknown'
        )
      )
  );
  comment on table credential_vault_credential_status_enm is
    'credential_vault_credential_status_enm is an enumeration table for the status of vault credentials. '
    'It contains rows for representing the active, revoke, revoked, expired, and unknown statuses.';

  insert into credential_vault_credential_status_enm (name)
  values
    ('active'),
    ('revoke'),
    ('revoked'),
    ('expired'),
    ('unknown');

  create table credential_vault_credential (
    public_id wt_public_id primary key,
    library_id wt_public_id not null
      constraint credential_vault_library_fkey
        references credential_vault_library (public_id)
        on delete cascade
        on update cascade,
    session_id wt_public_id not null
      constraint session_fkey
        references session (public_id)
        on delete cascade
        on update cascade,
    token_hmac bytea not null
      constraint credential_vault_token_fkey
        references credential_vault_token (token_hmac)
        on delete cascade
        on update cascade,
    create_time wt_timestamp,
    update_time wt_timestamp,
    version wt_version,
    external_id wt_sentinel not null,
    last_renewal_time timestamp with time zone not null,
    expiration_time timestamp with time zone not null
      constraint last_renewal_time_must_be_before_expiration_time
        check(last_renewal_time < expiration_time),
    is_renewable boolean not null,
    status text not null
      constraint credential_vault_credential_status_enm_fkey
        references credential_vault_credential_status_enm (name)
        on delete restrict
        on update cascade,
    constraint credential_dynamic_fkey
      foreign key (library_id, public_id)
      references credential_dynamic (library_id, public_id)
      on delete cascade
      on update cascade,
    constraint credential_vault_credential_library_id_public_id_uq
      unique(library_id, public_id)
  );
  comment on table credential_vault_credential is
    'credential_vault_credential is a table where each row contains the lease information for a single Vault secret retrieved from a vault credential library for a session.';

  create trigger update_version_column after update on credential_vault_credential
    for each row execute procedure update_version_column();

  create trigger update_time_column before update on credential_vault_credential
    for each row execute procedure update_time_column();

  create trigger default_create_time_column before insert on credential_vault_credential
    for each row execute procedure default_create_time();

  create trigger immutable_columns before update on credential_vault_credential
    for each row execute procedure immutable_columns('external_id', 'library_id','session_id', 'create_time');

  create trigger insert_credential_dynamic_subtype before insert on credential_vault_credential
    for each row execute procedure insert_credential_dynamic_subtype();

  create trigger delete_credential_dynamic_subtype after delete on credential_vault_credential
    for each row execute procedure delete_credential_dynamic_subtype();

  insert into oplog_ticket (name, version)
  values
    ('credential_vault_store', 1),
    ('credential_vault_library', 1),
    ('credential_vault_credential', 1) ;

     create view credential_vault_store_private as
     with
     active_tokens as (
        select token_hmac,
               token, -- encrypted
               store_id,
               create_time,
               update_time,
               last_renewal_time,
               expiration_time,
               -- renewal time is the midpoint between the last renewal time and the expiration time
               last_renewal_time + (expiration_time - last_renewal_time) / 2 as renewal_time,
               key_id,
               status
          from credential_vault_token
         where status in ('current', 'maintaining')
     )
     select store.public_id           as public_id,
            store.scope_id            as scope_id,
            store.name                as name,
            store.description         as description,
            store.create_time         as create_time,
            store.update_time         as update_time,
            store.version             as version,
            store.vault_address       as vault_address,
            store.namespace           as namespace,
            store.ca_cert             as ca_cert,
            store.tls_server_name     as tls_server_name,
            store.tls_skip_verify     as tls_skip_verify,
            store.public_id           as store_id,
            token.token_hmac          as token_hmac,
            token.token               as ct_token, -- encrypted
            token.create_time         as token_create_time,
            token.update_time         as token_update_time,
            token.last_renewal_time   as token_last_renewal_time,
            token.expiration_time     as token_expiration_time,
            token.renewal_time        as token_renewal_time,
            token.key_id              as token_key_id,
            token.status              as token_status,
            cert.certificate          as client_cert,
            cert.certificate_key      as ct_client_key, -- encrypted
            cert.certificate_key_hmac as client_cert_key_hmac,
            cert.key_id               as client_key_id
       from credential_vault_store store
  left join active_tokens token
         on store.public_id = token.store_id
  left join credential_vault_client_certificate cert
         on store.public_id = cert.store_id;
  comment on view credential_vault_store_private is
    'credential_vault_store_private is a view where each row contains a credential store and the credential store''s data needed to connect to Vault. '
    'The view returns a separate row for each current and maintaining token, maintaining tokens should only be used for token/credential renewal and revocation. '
    'Each row may contain encrypted data. This view should not be used to retrieve data which will be returned external to boundary.';

     create view credential_vault_store_public as
     select public_id,
            scope_id,
            name,
            description,
            create_time,
            update_time,
            version,
            vault_address,
            namespace,
            ca_cert,
            tls_server_name,
            tls_skip_verify,
            token_hmac,
            token_create_time,
            token_update_time,
            token_last_renewal_time,
            token_expiration_time,
            client_cert,
            client_cert_key_hmac
       from credential_vault_store_private
      where token_status = 'current';
  comment on view credential_vault_store_public is
    'credential_vault_store_public is a view where each row contains a credential store. '
    'No encrypted data is returned. This view can be used to retrieve data which will be returned external to boundary.';

     create view credential_vault_library_private as
     select library.public_id         as public_id,
            library.store_id          as store_id,
            library.name              as name,
            library.description       as description,
            library.create_time       as create_time,
            library.update_time       as update_time,
            library.version           as version,
            library.vault_path        as vault_path,
            library.http_method       as http_method,
            library.http_request_body as http_request_body,
            store.scope_id            as scope_id,
            store.vault_address       as vault_address,
            store.namespace           as namespace,
            store.ca_cert             as ca_cert,
            store.tls_server_name     as tls_server_name,
            store.tls_skip_verify     as tls_skip_verify,
            store.token_hmac          as token_hmac,
            store.ct_token            as ct_token, -- encrypted
            store.token_key_id        as token_key_id,
            store.client_cert         as client_cert,
            store.ct_client_key       as ct_client_key, -- encrypted
            store.client_key_id       as client_key_id
       from credential_vault_library library
       join credential_vault_store_private store
         on library.store_id = store.public_id
        and store.token_status = 'current';
  comment on view credential_vault_library_private is
    'credential_vault_library_private is a view where each row contains a credential library and the credential library''s data needed to connect to Vault. '
    'Each row may contain encrypted data. This view should not be used to retrieve data which will be returned external to boundary.';

     create view credential_vault_credential_private as
     select credential.public_id         as public_id,
            credential.library_id        as library_id,
            credential.session_id        as session_id,
            credential.create_time       as create_time,
            credential.update_time       as update_time,
            credential.version           as version,
            credential.external_id       as external_id,
            credential.last_renewal_time as last_renewal_time,
            credential.expiration_time   as expiration_time,
            credential.is_renewable      as is_renewable,
            credential.status            as status,
            credential.last_renewal_time + (credential.expiration_time - credential.last_renewal_time) / 2 as renewal_time,
            token.token_hmac             as token_hmac,
            token.token                  as ct_token, -- encrypted
            token.create_time            as token_create_time,
            token.update_time            as token_update_time,
            token.last_renewal_time      as token_last_renewal_time,
            token.expiration_time        as token_expiration_time,
            token.key_id                 as token_key_id,
            token.status                 as token_status,
            store.scope_id               as scope_id,
            store.vault_address          as vault_address,
            store.namespace              as namespace,
            store.ca_cert                as ca_cert,
            store.tls_server_name        as tls_server_name,
            store.tls_skip_verify        as tls_skip_verify,
            cert.certificate             as client_cert,
            cert.certificate_key         as ct_client_key, -- encrypted
            cert.certificate_key_hmac    as client_cert_key_hmac,
            cert.key_id                  as client_key_id
       from credential_vault_credential credential
       join credential_vault_token token
         on credential.token_hmac = token.token_hmac
       join credential_vault_store store
         on token.store_id = store.public_id
  left join credential_vault_client_certificate cert
         on store.public_id = cert.store_id
      where credential.expiration_time != 'infinity'::date;
  comment on view credential_vault_credential_private is
    'credential_vault_credential_private is a view where each row contains a credential, '
    'the vault token used to issue the credential, and the credential store data needed to connect to Vault. '
    'Each row may contain encrypted data. This view should not be used to retrieve data which will be returned external to boundary.';

commit;
