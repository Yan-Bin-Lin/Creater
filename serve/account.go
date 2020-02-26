package serve

/*
   salt := make([]byte, PW_SALT_BYTES)
   _, err := io.ReadFull(rand.Reader, salt)
   if err != nil {
       log.Fatal(err)
   }

   hash, err := scrypt.Key([]byte(password), salt, 1<<14, 8, 1, PW_HASH_BYTES)
   if err != nil {
       log.Fatal(err)
   }
*/
