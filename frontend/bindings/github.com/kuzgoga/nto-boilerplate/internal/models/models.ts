// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import {Create as $Create} from "@wailsio/runtime";

export class DryMode {
    "Id": number;
    "Name": string;
    "WetMaterialsId": number;
    "WetMaterials": WoodSpec;
    "HumidityPercent": number;
    "ProcentYsyshki": number;

    /** Creates a new DryMode instance. */
    constructor($$source: Partial<DryMode> = {}) {
        if (!("Id" in $$source)) {
            this["Id"] = 0;
        }
        if (!("Name" in $$source)) {
            this["Name"] = "";
        }
        if (!("WetMaterialsId" in $$source)) {
            this["WetMaterialsId"] = 0;
        }
        if (!("WetMaterials" in $$source)) {
            this["WetMaterials"] = (new WoodSpec());
        }
        if (!("HumidityPercent" in $$source)) {
            this["HumidityPercent"] = 0;
        }
        if (!("ProcentYsyshki" in $$source)) {
            this["ProcentYsyshki"] = 0;
        }

        Object.assign(this, $$source);
    }

    /**
     * Creates a new DryMode instance from a string or object.
     */
    static createFrom($$source: any = {}): DryMode {
        const $$createField3_0 = $$createType0;
        let $$parsedSource = typeof $$source === 'string' ? JSON.parse($$source) : $$source;
        if ("WetMaterials" in $$parsedSource) {
            $$parsedSource["WetMaterials"] = $$createField3_0($$parsedSource["WetMaterials"]);
        }
        return new DryMode($$parsedSource as Partial<DryMode>);
    }
}

export class Exporter {
    "Id": number;
    "Name": string;
    "Description": string;
    "Receivers": Receiver[];

    /** Creates a new Exporter instance. */
    constructor($$source: Partial<Exporter> = {}) {
        if (!("Id" in $$source)) {
            this["Id"] = 0;
        }
        if (!("Name" in $$source)) {
            this["Name"] = "";
        }
        if (!("Description" in $$source)) {
            this["Description"] = "";
        }
        if (!("Receivers" in $$source)) {
            this["Receivers"] = [];
        }

        Object.assign(this, $$source);
    }

    /**
     * Creates a new Exporter instance from a string or object.
     */
    static createFrom($$source: any = {}): Exporter {
        const $$createField3_0 = $$createType2;
        let $$parsedSource = typeof $$source === 'string' ? JSON.parse($$source) : $$source;
        if ("Receivers" in $$parsedSource) {
            $$parsedSource["Receivers"] = $$createField3_0($$parsedSource["Receivers"]);
        }
        return new Exporter($$parsedSource as Partial<Exporter>);
    }
}

export class Postav {
    "Id": number;
    "Name": string;
    "CenterDoskaSpecId": number;
    "CenterDoskaSpec": WoodSpec;
    "CenterOutPercent": number;
    "BacksideDoskaSpecId": number;
    "BacksideDoskaSpec": WoodSpec;
    "BacksideOutPercent": number;
    "OpilkiSpecId": number;
    "OpilkiSpec": WoodSpec;
    "OpilkiOutPercent": number;

    /** Creates a new Postav instance. */
    constructor($$source: Partial<Postav> = {}) {
        if (!("Id" in $$source)) {
            this["Id"] = 0;
        }
        if (!("Name" in $$source)) {
            this["Name"] = "";
        }
        if (!("CenterDoskaSpecId" in $$source)) {
            this["CenterDoskaSpecId"] = 0;
        }
        if (!("CenterDoskaSpec" in $$source)) {
            this["CenterDoskaSpec"] = (new WoodSpec());
        }
        if (!("CenterOutPercent" in $$source)) {
            this["CenterOutPercent"] = 0;
        }
        if (!("BacksideDoskaSpecId" in $$source)) {
            this["BacksideDoskaSpecId"] = 0;
        }
        if (!("BacksideDoskaSpec" in $$source)) {
            this["BacksideDoskaSpec"] = (new WoodSpec());
        }
        if (!("BacksideOutPercent" in $$source)) {
            this["BacksideOutPercent"] = 0;
        }
        if (!("OpilkiSpecId" in $$source)) {
            this["OpilkiSpecId"] = 0;
        }
        if (!("OpilkiSpec" in $$source)) {
            this["OpilkiSpec"] = (new WoodSpec());
        }
        if (!("OpilkiOutPercent" in $$source)) {
            this["OpilkiOutPercent"] = 0;
        }

        Object.assign(this, $$source);
    }

    /**
     * Creates a new Postav instance from a string or object.
     */
    static createFrom($$source: any = {}): Postav {
        const $$createField3_0 = $$createType0;
        const $$createField6_0 = $$createType0;
        const $$createField9_0 = $$createType0;
        let $$parsedSource = typeof $$source === 'string' ? JSON.parse($$source) : $$source;
        if ("CenterDoskaSpec" in $$parsedSource) {
            $$parsedSource["CenterDoskaSpec"] = $$createField3_0($$parsedSource["CenterDoskaSpec"]);
        }
        if ("BacksideDoskaSpec" in $$parsedSource) {
            $$parsedSource["BacksideDoskaSpec"] = $$createField6_0($$parsedSource["BacksideDoskaSpec"]);
        }
        if ("OpilkiSpec" in $$parsedSource) {
            $$parsedSource["OpilkiSpec"] = $$createField9_0($$parsedSource["OpilkiSpec"]);
        }
        return new Postav($$parsedSource as Partial<Postav>);
    }
}

export class Receiver {
    "Id": number;
    "ExporterId": number;
    "Exporter": Exporter;
    "ExporterQuantity": number;
    "FactoryQuantity": number;
    "Description": string;
    "Spec": WoodSpecType;
    "SpecId": number;
    "CreatedAt": number;

    /** Creates a new Receiver instance. */
    constructor($$source: Partial<Receiver> = {}) {
        if (!("Id" in $$source)) {
            this["Id"] = 0;
        }
        if (!("ExporterId" in $$source)) {
            this["ExporterId"] = 0;
        }
        if (!("Exporter" in $$source)) {
            this["Exporter"] = (new Exporter());
        }
        if (!("ExporterQuantity" in $$source)) {
            this["ExporterQuantity"] = 0;
        }
        if (!("FactoryQuantity" in $$source)) {
            this["FactoryQuantity"] = 0;
        }
        if (!("Description" in $$source)) {
            this["Description"] = "";
        }
        if (!("Spec" in $$source)) {
            this["Spec"] = (new WoodSpecType());
        }
        if (!("SpecId" in $$source)) {
            this["SpecId"] = 0;
        }
        if (!("CreatedAt" in $$source)) {
            this["CreatedAt"] = 0;
        }

        Object.assign(this, $$source);
    }

    /**
     * Creates a new Receiver instance from a string or object.
     */
    static createFrom($$source: any = {}): Receiver {
        const $$createField2_0 = $$createType3;
        const $$createField6_0 = $$createType4;
        let $$parsedSource = typeof $$source === 'string' ? JSON.parse($$source) : $$source;
        if ("Exporter" in $$parsedSource) {
            $$parsedSource["Exporter"] = $$createField2_0($$parsedSource["Exporter"]);
        }
        if ("Spec" in $$parsedSource) {
            $$parsedSource["Spec"] = $$createField6_0($$parsedSource["Spec"]);
        }
        return new Receiver($$parsedSource as Partial<Receiver>);
    }
}

export class SushkaResult {
    "Id": number;
    "Description": string;
    "MaterialSpecId": number;
    "MaterialSpec": WoodSpec;
    "MaterialQuantity": number;
    "ProductSpecId": number;
    "ProductSpec": WoodSpec;
    "ProductAmount": number;
    "WorkDate": number;

    /** Creates a new SushkaResult instance. */
    constructor($$source: Partial<SushkaResult> = {}) {
        if (!("Id" in $$source)) {
            this["Id"] = 0;
        }
        if (!("Description" in $$source)) {
            this["Description"] = "";
        }
        if (!("MaterialSpecId" in $$source)) {
            this["MaterialSpecId"] = 0;
        }
        if (!("MaterialSpec" in $$source)) {
            this["MaterialSpec"] = (new WoodSpec());
        }
        if (!("MaterialQuantity" in $$source)) {
            this["MaterialQuantity"] = 0;
        }
        if (!("ProductSpecId" in $$source)) {
            this["ProductSpecId"] = 0;
        }
        if (!("ProductSpec" in $$source)) {
            this["ProductSpec"] = (new WoodSpec());
        }
        if (!("ProductAmount" in $$source)) {
            this["ProductAmount"] = 0;
        }
        if (!("WorkDate" in $$source)) {
            this["WorkDate"] = 0;
        }

        Object.assign(this, $$source);
    }

    /**
     * Creates a new SushkaResult instance from a string or object.
     */
    static createFrom($$source: any = {}): SushkaResult {
        const $$createField3_0 = $$createType0;
        const $$createField6_0 = $$createType0;
        let $$parsedSource = typeof $$source === 'string' ? JSON.parse($$source) : $$source;
        if ("MaterialSpec" in $$parsedSource) {
            $$parsedSource["MaterialSpec"] = $$createField3_0($$parsedSource["MaterialSpec"]);
        }
        if ("ProductSpec" in $$parsedSource) {
            $$parsedSource["ProductSpec"] = $$createField6_0($$parsedSource["ProductSpec"]);
        }
        return new SushkaResult($$parsedSource as Partial<SushkaResult>);
    }
}

export class WoodSpec {
    "Id": number;
    "WoodSpecTypeId": number;
    "WoodSpecType": WoodSpecType;
    "Poroda": string;
    "Sort": string;
    "DlinaBrevna": string;
    "DiameterBrevna": string;
    "SechenieDoski": string;
    "ProcentVlajnosti": number;
    "DiameterGranyl": string;

    /** Creates a new WoodSpec instance. */
    constructor($$source: Partial<WoodSpec> = {}) {
        if (!("Id" in $$source)) {
            this["Id"] = 0;
        }
        if (!("WoodSpecTypeId" in $$source)) {
            this["WoodSpecTypeId"] = 0;
        }
        if (!("WoodSpecType" in $$source)) {
            this["WoodSpecType"] = (new WoodSpecType());
        }
        if (!("Poroda" in $$source)) {
            this["Poroda"] = "";
        }
        if (!("Sort" in $$source)) {
            this["Sort"] = "";
        }
        if (!("DlinaBrevna" in $$source)) {
            this["DlinaBrevna"] = "";
        }
        if (!("DiameterBrevna" in $$source)) {
            this["DiameterBrevna"] = "";
        }
        if (!("SechenieDoski" in $$source)) {
            this["SechenieDoski"] = "";
        }
        if (!("ProcentVlajnosti" in $$source)) {
            this["ProcentVlajnosti"] = 0;
        }
        if (!("DiameterGranyl" in $$source)) {
            this["DiameterGranyl"] = "";
        }

        Object.assign(this, $$source);
    }

    /**
     * Creates a new WoodSpec instance from a string or object.
     */
    static createFrom($$source: any = {}): WoodSpec {
        const $$createField2_0 = $$createType4;
        let $$parsedSource = typeof $$source === 'string' ? JSON.parse($$source) : $$source;
        if ("WoodSpecType" in $$parsedSource) {
            $$parsedSource["WoodSpecType"] = $$createField2_0($$parsedSource["WoodSpecType"]);
        }
        return new WoodSpec($$parsedSource as Partial<WoodSpec>);
    }
}

export class WoodSpecType {
    "Id": number;
    "Name": string;
    "Receivers": Receiver[];

    /** Creates a new WoodSpecType instance. */
    constructor($$source: Partial<WoodSpecType> = {}) {
        if (!("Id" in $$source)) {
            this["Id"] = 0;
        }
        if (!("Name" in $$source)) {
            this["Name"] = "";
        }
        if (!("Receivers" in $$source)) {
            this["Receivers"] = [];
        }

        Object.assign(this, $$source);
    }

    /**
     * Creates a new WoodSpecType instance from a string or object.
     */
    static createFrom($$source: any = {}): WoodSpecType {
        const $$createField2_0 = $$createType2;
        let $$parsedSource = typeof $$source === 'string' ? JSON.parse($$source) : $$source;
        if ("Receivers" in $$parsedSource) {
            $$parsedSource["Receivers"] = $$createField2_0($$parsedSource["Receivers"]);
        }
        return new WoodSpecType($$parsedSource as Partial<WoodSpecType>);
    }
}

export class WorkResult {
    "Id": number;
    "MaterialId": number;
    "Material": WoodSpecType;
    "MaterialQuantity": number;
    "ProductSpecId": number;
    "ProductSpec": WoodSpec;
    "WorkDate": number;

    /** Creates a new WorkResult instance. */
    constructor($$source: Partial<WorkResult> = {}) {
        if (!("Id" in $$source)) {
            this["Id"] = 0;
        }
        if (!("MaterialId" in $$source)) {
            this["MaterialId"] = 0;
        }
        if (!("Material" in $$source)) {
            this["Material"] = (new WoodSpecType());
        }
        if (!("MaterialQuantity" in $$source)) {
            this["MaterialQuantity"] = 0;
        }
        if (!("ProductSpecId" in $$source)) {
            this["ProductSpecId"] = 0;
        }
        if (!("ProductSpec" in $$source)) {
            this["ProductSpec"] = (new WoodSpec());
        }
        if (!("WorkDate" in $$source)) {
            this["WorkDate"] = 0;
        }

        Object.assign(this, $$source);
    }

    /**
     * Creates a new WorkResult instance from a string or object.
     */
    static createFrom($$source: any = {}): WorkResult {
        const $$createField2_0 = $$createType4;
        const $$createField5_0 = $$createType0;
        let $$parsedSource = typeof $$source === 'string' ? JSON.parse($$source) : $$source;
        if ("Material" in $$parsedSource) {
            $$parsedSource["Material"] = $$createField2_0($$parsedSource["Material"]);
        }
        if ("ProductSpec" in $$parsedSource) {
            $$parsedSource["ProductSpec"] = $$createField5_0($$parsedSource["ProductSpec"]);
        }
        return new WorkResult($$parsedSource as Partial<WorkResult>);
    }
}

// Private type creation functions
const $$createType0 = WoodSpec.createFrom;
const $$createType1 = Receiver.createFrom;
const $$createType2 = $Create.Array($$createType1);
const $$createType3 = Exporter.createFrom;
const $$createType4 = WoodSpecType.createFrom;
