package solace.coverage.metrics;

import java.util.Collection;


import io.swagger.v3.oas.models.Operation;
import io.swagger.v3.oas.models.parameters.Parameter;
import solace.coverage.base.AbstractMetric;
import solace.coverage.base.MetricMetadata;

public class OperationMetric extends AbstractMetric<Operation> {

    public OperationMetric(Operation entity, Collection<Parameter> defaultParameters) {
        super(entity);
        //TODO Merge default params
    }

    @Override
    public float getCoverage() {
        // TODO Auto-generated method stub
        return 0;
    }

    @Override
    public void Reset() {
        // TODO Auto-generated method stub

    }

    @Override
    public MetricMetadata getMetadata() {
        // TODO Auto-generated method stub
        return null;
    }

}