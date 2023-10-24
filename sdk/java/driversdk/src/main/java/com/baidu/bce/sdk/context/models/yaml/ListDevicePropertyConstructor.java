package com.baidu.bce.sdk.context.models.yaml;

import com.baidu.bce.sdk.context.models.device.DeviceProperty;
import org.yaml.snakeyaml.LoaderOptions;
import org.yaml.snakeyaml.constructor.Constructor;
import org.yaml.snakeyaml.nodes.MappingNode;
import org.yaml.snakeyaml.nodes.Node;
import org.yaml.snakeyaml.nodes.SequenceNode;
import org.yaml.snakeyaml.nodes.Tag;

import java.util.stream.Collectors;

public class ListDevicePropertyConstructor extends Constructor {
    private static final String MODEL_YAML_CONSTRUCTOR_NAME = ListDevicePropertyConstructor.class.getSimpleName();

    public ListDevicePropertyConstructor(LoaderOptions loadingConfig) {
        super(loadingConfig);
        this.rootTag = new Tag(MODEL_YAML_CONSTRUCTOR_NAME);
    }

    @Override
    protected Object constructObject(Node node) {
        if (MODEL_YAML_CONSTRUCTOR_NAME.equals(node.getTag().getValue())
                && node instanceof MappingNode mNode) {
            return mNode.getValue().stream().collect(
                    Collectors.toMap(
                            t -> super.constructObject(t.getKeyNode()),
                            t -> {
                                Node child = t.getValueNode();
                                if (child instanceof SequenceNode sNode){
                                    sNode.setListType(DeviceProperty.class);
                                }
                                return super.constructObject(child);
                            }
                    )
            );
        }
        return super.constructObject(node);
    }
}
