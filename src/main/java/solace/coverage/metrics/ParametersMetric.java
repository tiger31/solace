package solace.coverage.metrics;

import java.util.Collection;

import io.swagger.v3.oas.models.parameters.Parameter;
import solace.coverage.base.AbstractCollectionMetric;

public class ParametersMetric extends AbstractCollectionMetric<Collection<Parameter>, Parameter> {
    private static final long serialVersionUID = 209347903780L;

    public ParametersMetric(Collection<Parameter> entity) {
        super(entity);
    }
    
}