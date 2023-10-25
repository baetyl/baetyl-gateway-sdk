package com.baidu.bce.sdk.context.models.yaml;

import com.baidu.bce.sdk.context.models.access.AccessTemplate;
import org.yaml.snakeyaml.LoaderOptions;
import org.yaml.snakeyaml.TypeDescription;
import org.yaml.snakeyaml.constructor.Constructor;
import org.yaml.snakeyaml.nodes.MappingNode;
import org.yaml.snakeyaml.nodes.Node;
import org.yaml.snakeyaml.nodes.Tag;

import java.util.stream.Collectors;

public class AccessTemplateConstructor extends Constructor {
    private static final String ACCESS_TEMPLATE_YAML_CONSTRUCTOR_NAME = AccessTemplateConstructor.class.getSimpleName();

    private final TypeDescription itemType = new TypeDescription(AccessTemplate.class);

    public AccessTemplateConstructor(LoaderOptions loadingConfig) {
        super(loadingConfig);
        this.rootTag = new Tag(ACCESS_TEMPLATE_YAML_CONSTRUCTOR_NAME);
        this.addTypeDescription(itemType);
    }

    @Override
    protected Object constructObject(Node node) {
        if (ACCESS_TEMPLATE_YAML_CONSTRUCTOR_NAME.equals(node.getTag().getValue())
                && node instanceof MappingNode mNode) {
            return mNode.getValue().stream().collect(
                    Collectors.toMap(
                            t -> super.constructObject(t.getKeyNode()),
                            t -> {
                                Node child = t.getValueNode();
                                child.setType(itemType.getType());
                                return super.constructObject(child);
                            }
                    )
            );
        }
        return super.constructObject(node);
    }
}
