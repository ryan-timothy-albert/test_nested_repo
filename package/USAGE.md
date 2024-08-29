<!-- Start SDK Example Usage [usage] -->
```typescript
import { Petstore } from "total-test";

const petstore = new Petstore();

async function run() {
    const result = await petstore.pets.listPets({});

    // Handle the result
    console.log(result);
}

run();

```
<!-- End SDK Example Usage [usage] -->