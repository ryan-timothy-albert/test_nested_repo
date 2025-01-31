<!-- Start SDK Example Usage [usage] -->
```typescript
import { Pets } from "ryan-total-test-act";

const pets = new Pets();

async function run() {
  const result = await pets.pets.listPets({});

  // Handle the result
  console.log(result);
}

run();

```
<!-- End SDK Example Usage [usage] -->