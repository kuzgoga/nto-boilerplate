// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT

/**
 * TODO: implement service and migrator
 * @module
 */

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import {Call as $Call, Create as $Create} from "@wailsio/runtime";

export function Greet(name: string): Promise<string> & { cancel(): void } {
    let $resultPromise = $Call.ByID(4216357642, name) as any;
    return $resultPromise;
}
