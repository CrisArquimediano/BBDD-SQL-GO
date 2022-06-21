\c tarjetas-de-credito

ALTER TABLE consumo DROP CONSTRAINT consumo_nrocomercio_fk;
ALTER TABLE consumo DROP CONSTRAINT consumo_nrotarjeta_fk;
ALTER TABLE alerta DROP CONSTRAINT alerta_nrorechazo_fk;
ALTER TABLE alerta DROP CONSTRAINT alerta_nrotarjeta_fk;
ALTER TABLE detalle DROP CONSTRAINT detalle_nroresumen_fk;
ALTER TABLE cabecera DROP CONSTRAINT cabecera_nrotarjeta_fk;
ALTER TABLE rechazo DROP CONSTRAINT rechazo_nrocomercio_fk;
ALTER TABLE rechazo DROP CONSTRAINT rechazo_nrotarjeta_fk;
ALTER TABLE compra DROP CONSTRAINT compra_nrocomercio_fk;
ALTER TABLE compra DROP CONSTRAINT compra_nrotarjeta_fk;
ALTER TABLE tarjeta DROP CONSTRAINT tarjeta_nrocliente_fk;

ALTER TABLE alerta DROP CONSTRAINT alerta_pk;
ALTER TABLE detalle DROP CONSTRAINT detalle_pk;
ALTER TABLE cabecera DROP CONSTRAINT cabecera_pk;
ALTER TABLE cierre DROP CONSTRAINT cierre_pk;
ALTER TABLE rechazo DROP CONSTRAINT rechazo_pk;
ALTER TABLE compra DROP CONSTRAINT compra_pk;
ALTER TABLE comercio DROP CONSTRAINT comercio_pk;
ALTER TABLE tarjeta DROP CONSTRAINT tarjeta_pk;
ALTER TABLE cliente DROP CONSTRAINT cliente_pk;

\c postgres