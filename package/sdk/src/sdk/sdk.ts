/*
 * Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.
 */

import { ClientSDK } from "../lib/sdks.js";
import { Pets } from "./pets.js";

export class SDK extends ClientSDK {
    private _pets?: Pets;
    get pets(): Pets {
        return (this._pets ??= new Pets(this.options$));
    }
}
